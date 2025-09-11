package exporter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jdpx/airthings-prometheus-exporter/internal/airthings"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const defaultAirthingsTokenURL = "https://accounts-api.airthings.com/v1/token"

// Config holds exporter configuration; fields are annotated for envconfig.
type Config struct {
	// General
	ListenAddr    string        `envconfig:"LISTEN_ADDR" default:":9000"`
	Unit          string        `envconfig:"UNIT" default:"metric"`
	PollInterval  time.Duration `envconfig:"POLL_INTERVAL" default:"60s"`
	IncludeSerial string        `envconfig:"INCLUDE_SERIALS"`
	AccountID     string        `envconfig:"ACCOUNT_ID"`
	LogLevel      string        `envconfig:"LOG_LEVEL" default:"info"`
	LogFormat     string        `envconfig:"LOG_FORMAT" default:"json"`

	// Authentication
	AirthingsToken string `envconfig:"AIRTHINGS_TOKEN"`

	// OAuth2 client credentials (optional, used if Token is empty)
	OAuthClientID     string `envconfig:"AIRTHINGS_CLIENT_ID,required"`
	OAuthClientSecret string `envconfig:"AIRTHINGS_CLIENT_SECRET,required"`
	OAuthTokenURL     string `envconfig:"AIRTHINGS_TOKEN_URL" default:"https://accounts-api.airthings.com/v1/token"`
	OAuthScope        string `envconfig:"AIRTHINGS_SCOPE" default:"read:device:current_values"`
	OAuthAudience     string `envconfig:"AIRTHINGS_AUDIENCE"`
}

type Exporter struct {
	cfg      Config
	client   *airthings.APIClient
	rt       *rateLimitCaptureRoundTripper
	registry *prometheus.Registry

	// metrics
	sensorGauge          *prometheus.GaugeVec
	batteryGauge         *prometheus.GaugeVec
	lastSuccessTimestamp prometheus.Gauge
	scrapeDuration       prometheus.Histogram
	apiLimit             prometheus.Gauge
	apiRemaining         prometheus.Gauge
	rateLimitedTotal     prometheus.Counter
	requestErrorsTotal   prometheus.Counter

	mu     sync.RWMutex
	devMap map[string]deviceInfo
}

type deviceInfo struct {
	Name string
	Type string
}

func New(cfg Config) (*Exporter, error) {
	if cfg.Unit != "metric" && cfg.Unit != "imperial" {
		return nil, fmt.Errorf("invalid UNIT: %s", cfg.Unit)
	}

	var httpClient *http.Client
	var rlCapture = &rateLimitCaptureRoundTripper{underlying: http.DefaultTransport}

	if strings.TrimSpace(cfg.AirthingsToken) != "" {
		// Static bearer token mode
		httpClient = &http.Client{Transport: &bearerRoundTripper{
			token:      cfg.AirthingsToken,
			underlying: rlCapture,
		}}
	} else if cfg.OAuthClientID != "" && cfg.OAuthClientSecret != "" {
		// OAuth2 client credentials mode
		if strings.TrimSpace(cfg.OAuthTokenURL) == "" {
			cfg.OAuthTokenURL = defaultAirthingsTokenURL
		}
		cc := &clientcredentials.Config{
			ClientID:     cfg.OAuthClientID,
			ClientSecret: cfg.OAuthClientSecret,
			TokenURL:     cfg.OAuthTokenURL,
		}
		if cfg.OAuthScope != "" {
			cc.Scopes = []string{cfg.OAuthScope}
		}
		if cfg.OAuthAudience != "" {
			if cc.EndpointParams == nil {
				cc.EndpointParams = url.Values{}
			}
			cc.EndpointParams.Set("audience", cfg.OAuthAudience)
		}
		ts := cc.TokenSource(context.Background())
		oauthTransport := &oauth2.Transport{Source: ts, Base: rlCapture}
		httpClient = &http.Client{Transport: oauthTransport, Timeout: 30 * time.Second}
	} else {
		return nil, errors.New("no AIRTHINGS_TOKEN provided and OAuth client credentials not configured")
	}

	apiCfg := airthings.NewConfiguration()
	apiCfg.HTTPClient = httpClient
	client := airthings.NewAPIClient(apiCfg)

	reg := prometheus.NewRegistry()
	e := &Exporter{
		cfg:      cfg,
		client:   client,
		rt:       rlCapture,
		registry: reg,
		devMap:   map[string]deviceInfo{},
	}

	e.sensorGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "airthings_sensor_value",
		Help: "Airthings sensor value",
	}, []string{"serial_number", "sensor_type", "unit", "device_name", "device_type"})

	e.batteryGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "airthings_battery_percentage",
		Help: "Airthings battery percentage",
	}, []string{"device_name", "serial_number"})

	e.lastSuccessTimestamp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airthings_last_scrape_success_timestamp_seconds",
		Help: "Unix timestamp of last successful poll",
	})

	e.scrapeDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "airthings_scrape_duration_seconds",
		Help:    "Duration of background poll",
		Buckets: prometheus.DefBuckets,
	})

	e.apiLimit = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airthings_api_limit",
		Help: "API hourly request limit",
	})

	e.apiRemaining = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "airthings_api_remaining",
		Help: "Remaining requests in current window",
	})

	e.rateLimitedTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "airthings_rate_limited_total",
		Help: "Number of 429 responses",
	})

	e.requestErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "airthings_request_errors_total",
		Help: "Number of non-429 request errors",
	})

	reg.MustRegister(
		e.sensorGauge, e.batteryGauge,
		e.lastSuccessTimestamp, e.scrapeDuration,
		e.apiLimit, e.apiRemaining,
		e.rateLimitedTotal, e.requestErrorsTotal,
	)

	// Also expose on default registry
	prometheus.MustRegister(
		e.sensorGauge, e.batteryGauge,
		e.lastSuccessTimestamp, e.scrapeDuration,
		e.apiLimit, e.apiRemaining,
		e.rateLimitedTotal, e.requestErrorsTotal,
	)

	return e, nil
}

// bearerRoundTripper injects static Bearer token
type bearerRoundTripper struct {
	underlying http.RoundTripper
	token      string
}

func (b *bearerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := req.Clone(req.Context())
	req2.Header.Set("Authorization", "Bearer "+b.token)
	return b.underlying.RoundTrip(req2)
}

// rateLimitCaptureRoundTripper captures rate limit headers
type rateLimitCaptureRoundTripper struct {
	underlying http.RoundTripper

	mu        sync.Mutex
	limit     float64
	remaining float64
	retryAt   time.Time
}

func (a *rateLimitCaptureRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := a.underlying.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	// Capture headers
	a.mu.Lock()
	if v := resp.Header.Get("X-RateLimit-Limit"); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			a.limit = f
		}
	}
	if v := resp.Header.Get("X-RateLimit-Remaining"); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			a.remaining = f
		}
	}
	// Retry-After can be seconds or date
	if v := resp.Header.Get("X-RateLimit-Retry-After"); v != "" {
		if secs, err := strconv.Atoi(v); err == nil {
			a.retryAt = time.Now().Add(time.Duration(secs) * time.Second)
		} else if t, err := http.ParseTime(v); err == nil {
			a.retryAt = t
		}
	}
	a.mu.Unlock()
	return resp, nil
}

func (e *Exporter) Run(ctx context.Context) {
	for {
		e.pollOnce(ctx)
		// adjust sleep for rate limit if needed
		nextSleep := e.cfg.PollInterval

		e.rt.mu.Lock()
		limit := e.rt.limit
		remaining := e.rt.remaining
		retryAt := e.rt.retryAt
		e.rt.mu.Unlock()

		if limit > 0 {
			e.apiLimit.Set(limit)
		}
		if remaining >= 0 {
			e.apiRemaining.Set(remaining)
		}
		if retryAt.After(time.Now()) {
			delay := time.Until(retryAt)
			if delay > nextSleep {
				nextSleep = delay
			}
		} else if remaining >= 0 && remaining < 2 {
			// throttle if remainder is very low
			nextSleep = maxDuration(nextSleep*2, 2*time.Minute)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(nextSleep):
		}
	}
}

func (e *Exporter) pollOnce(ctx context.Context) {
	start := time.Now()
	defer func() { e.scrapeDuration.Observe(time.Since(start).Seconds()) }()

	accountID := e.cfg.AccountID
	if accountID == "" {
		id, err := e.discoverAccount(ctx)
		if err != nil {
			e.requestErrorsTotal.Inc()
			log.WithError(err).Warn("discover account failed")
			return
		}
		accountID = id
	}

	if err := e.refreshDevices(ctx, accountID); err != nil {
		e.requestErrorsTotal.Inc()
		log.WithError(err).Warn("refresh devices failed")
		// continue; still try to fetch sensors
	}

	serialFilter := setFromCSV(e.cfg.IncludeSerial)

	page := int32(1)
	for {
		req := e.client.SensorAPI.GetMultipleSensors(ctx, accountID).
			PageNumber(page).
			Unit(e.cfg.Unit)
		if len(serialFilter) > 0 {
			req = req.Sn(sliceFromSet(serialFilter))
		}
		resp, httpResp, err := req.Execute()
		if err != nil {
			if httpResp != nil && httpResp.StatusCode == http.StatusTooManyRequests {
				e.rateLimitedTotal.Inc()
				return
			}
			e.requestErrorsTotal.Inc()
			log.WithError(err).Warn("get sensors request failed")
			return
		}
		if resp == nil {
			e.requestErrorsTotal.Inc()
			log.Warn("nil sensors response")
			return
		}

		for _, sr := range resp.Results {
			sn := sr.GetSerialNumber()
			var dev deviceInfo
			e.mu.RLock()
			dev = e.devMap[sn]
			e.mu.RUnlock()
			if serialFilter != nil && !serialFilter[sn] {
				continue
			}

			if bp, ok := sr.GetBatteryPercentageOk(); ok && bp != nil {
				e.batteryGauge.WithLabelValues(sn).Set(float64(*bp))
			}
			for _, s := range sr.GetSensors() {
				unit := s.GetUnit()
				e.sensorGauge.WithLabelValues(
					sn,
					s.GetSensorType(),
					unit,
					dev.Name,
					dev.Type,
				).Set(s.GetValue())
			}
		}

		if resp.HasNext != nil && !*resp.HasNext {
			break
		}
		page++
	}

	e.lastSuccessTimestamp.Set(float64(time.Now().Unix()))
}

func (e *Exporter) refreshDevices(ctx context.Context, accountID string) error {
	resp, httpResp, err := e.client.DeviceAPI.GetDevices(ctx, accountID).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusTooManyRequests {
			e.rateLimitedTotal.Inc()
			return errors.New("rate limited")
		}
		return err
	}
	if resp == nil {
		return fmt.Errorf("devices response nil")
	}

	m := map[string]deviceInfo{}
	for _, d := range resp.Devices {
		m[d.GetSerialNumber()] = deviceInfo{
			Name: d.GetName(),
			Type: d.GetType(),
		}
	}

	e.mu.Lock()
	e.devMap = m
	e.mu.Unlock()
	return nil
}

func (e *Exporter) discoverAccount(ctx context.Context) (string, error) {
	resp, httpResp, err := e.client.AccountsAPI.GetAccountsIds(ctx).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == http.StatusTooManyRequests {
			e.rateLimitedTotal.Inc()
			return "", errors.New("rate limited")
		}
		return "", err
	}
	if resp == nil || len(resp.Accounts) == 0 {
		return "", errors.New("no accounts found")
	}
	return resp.Accounts[0].GetId(), nil
}

// Helpers

func setFromCSV(csv string) map[string]bool {
	if strings.TrimSpace(csv) == "" {
		return nil
	}
	out := map[string]bool{}
	for _, s := range strings.Split(csv, ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			out[s] = true
		}
	}
	return out
}

func sliceFromSet(m map[string]bool) []string {
	if m == nil {
		return nil
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func maxDuration(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}
