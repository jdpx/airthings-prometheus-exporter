package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jdpx/airthings-prometheus-exporter/internal/exporter"
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	var cfg exporter.Config
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}

	// Configure logging
	if cfg.LogFormat == "text" {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	} else {
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: time.RFC3339})
	}
	if cfg.LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if cfg.AirthingsToken == "" && (cfg.OAuthClientID == "" || cfg.OAuthClientSecret == "") {
		log.Fatal("either AIRTHINGS_TOKEN or OAuth client credentials (AIRTHINGS_CLIENT_ID/SECRET) must be provided")
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
