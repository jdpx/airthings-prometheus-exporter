package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jdpx/airthings-prometheus-exporter/internal/exporter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	level := os.Getenv("LOG_LEVEL")
	if level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	cfg := exporter.Config{
		Token:         os.Getenv("AIRTHINGS_TOKEN"),
		AccountID:     os.Getenv("ACCOUNT_ID"),
		Unit:          getOr("UNIT", "metric"),
		PollInterval:  getDurationOr("POLL_INTERVAL", 60*time.Second),
		IncludeSerial: os.Getenv("INCLUDE_SERIALS"),
		ListenAddr:    getOr("LISTEN_ADDR", ":9000"),
	}
	if cfg.Token == "" {
		log.Fatal("AIRTHINGS_TOKEN is required")
	}

	exp, err := exporter.New(cfg)
	if err != nil {
		log.WithError(err).Fatal("failed to init exporter")
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go exp.Run(ctx)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(200) })

	srv := &http.Server{
		Addr:              cfg.ListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.WithField("addr", cfg.ListenAddr).Info("listening")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("server error")
	}
}

func getOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func getDurationOr(k string, def time.Duration) time.Duration {
	if v := os.Getenv(k); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}
