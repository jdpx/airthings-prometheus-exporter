# Airthings Prometheus Exporter

A lightweight Prometheus exporter which periodically fetches sensor data from the Airthings for Consumer API and exposes them as Prometheus metrics. Designed to run in a small container and respect strict API rate limits by polling in the background and serving cached metrics.

## Features

- Background poller with cached metrics (no API calls during `/metrics` scrapes)
- Handles API rate limits and backs off using `X-RateLimit-*` headers
- Exposes rate limit state and error counters as metrics
- Configurable units, poll interval, device filtering
- Supports static Bearer token or OAuth2 Client Credentials
- Multi-arch container images published via GoReleaser

## How it works

- A background worker polls the Airthings API on a schedule, respecting rate limits and pagination.
- Results are cached in-memory and exposed on `/metrics` using Prometheus metrics.
- If `ACCOUNT_ID` is not provided, the exporter discovers the single account automatically.

Client code is generated from the API specification at `openapi.yaml` using `openapi-generator`.

## Authentication

You can authenticate either with a ready-to-use Bearer token or via OAuth2 Client Credentials.

- Static token (simplest):
  - `AIRTHINGS_TOKEN`: Provided token is sent as `Authorization: Bearer <token>`.
- OAuth2 Client Credentials (machine-to-machine), per Airthings auth docs ([Authorization docs](https://consumer-api-doc.airthings.com/docs/api/authorization)):
  - `AIRTHINGS_CLIENT_ID` (required if no token)
  - `AIRTHINGS_CLIENT_SECRET` (required if no token)
  - `AIRTHINGS_TOKEN_URL` (default: `https://accounts-api.airthings.com/v1/token`)
  - `AIRTHINGS_SCOPE` (default: `read:device:current_values`)
  - `AIRTHINGS_AUDIENCE` (optional, if your provider requires an audience claim)

If `AIRTHINGS_TOKEN` is unset and client credentials are configured, the exporter fetches and refreshes an access token automatically.

## Configuration (environment variables)

General:
- `LISTEN_ADDR` (default `:9000`): HTTP listen address
- `LOG_LEVEL` (default `info`): `info` or `debug`
- `LOG_FORMAT` (default `json`): `json` or `text`
- `UNIT` (default `metric`): `metric` or `imperial`
- `POLL_INTERVAL` (default `60s`): Background polling cadence
- `ACCOUNT_ID` (optional): If unset, auto-discovers the single account
- `INCLUDE_SERIALS` (optional): Comma-separated list of serials to include

Auth:
- `AIRTHINGS_TOKEN` (optional): Static Bearer token
- `AIRTHINGS_CLIENT_ID` / `AIRTHINGS_CLIENT_SECRET` (required if `AIRTHINGS_TOKEN` absent)
- `AIRTHINGS_TOKEN_URL` (default: `https://accounts-api.airthings.com/v1/token`)
- `AIRTHINGS_SCOPE` (default: `read:device:current_values`)
- `AIRTHINGS_AUDIENCE` (optional)

## Exposed metrics

- `airthings_sensor_value{serial_number, sensor_type, unit, device_name, device_type}`: Current sensor values
- `airthings_battery_percentage{serial_number}`: Battery percentage per device
- `airthings_last_scrape_success_timestamp_seconds`: Unix timestamp of last successful poll
- `airthings_scrape_duration_seconds`: Duration of the background poll
- `airthings_api_limit`: Current API hourly request limit
- `airthings_api_remaining`: Remaining requests in the current window
- `airthings_rate_limited_total`: Counter of `429` responses
- `airthings_request_errors_total`: Counter of non-429 request errors

Notes:
- Labels are kept intentionally small to avoid cardinality issues. If your fleet is large, consider using `INCLUDE_SERIALS`.

## Running locally

With Go installed (Bearer token example):

```bash
AIRTHINGS_TOKEN=your_token \
UNIT=metric \
POLL_INTERVAL=60s \
LISTEN_ADDR=":9000" \
INCLUDE_SERIALS="" \
LOG_LEVEL=info \
LOG_FORMAT=json \
go run ./cmd/airthings_exporter
```

OAuth2 client credentials example:

```bash
AIRTHINGS_CLIENT_ID=... \
AIRTHINGS_CLIENT_SECRET=... \
AIRTHINGS_TOKEN_URL=https://accounts-api.airthings.com/v1/token \
AIRTHINGS_SCOPE=read:device:current_values \
AIRTHINGS_AUDIENCE="" \
UNIT=metric \
POLL_INTERVAL=60s \
LOG_FORMAT=json \
go run ./cmd/airthings_exporter
```

Then visit `http://localhost:9000/metrics`.

## Docker

Build locally:

```bash
docker build -t ghcr.io/jdpx/airthings_prometheus_exporter:dev .
```

Run (Bearer token):

```bash
docker run --rm -p 9000:9000 \
  -e AIRTHINGS_TOKEN=your_token \
  -e UNIT=metric \
  -e LOG_FORMAT=json \
  ghcr.io/jdpx/airthings_prometheus_exporter:dev
```

Run (OAuth2 client credentials):

```bash
docker run --rm -p 9000:9000 \
  -e AIRTHINGS_CLIENT_ID=... \
  -e AIRTHINGS_CLIENT_SECRET=... \
  -e AIRTHINGS_TOKEN_URL=https://accounts-api.airthings.com/v1/token \
  -e AIRTHINGS_SCOPE=read:device:current_values \
  -e LOG_FORMAT=json \
  ghcr.io/jdpx/airthings_prometheus_exporter:dev
```

### docker-compose

A sample compose file is included; see below for all supported env vars.

```bash
docker compose up -d
```

Set env in your shell or a `.env` file.

## Prometheus scrape config

```yaml
- job_name: airthings
  static_configs:
    - targets: ["airthings-exporter:9000"]
```

## Development

- Go version: 1.24
- Linting: `golangci-lint` via CI
- Build locally: `go build ./...`
- Test: `go test ./...`

## Releases & CI

- Images are built and published to GHCR via GoReleaser on pushes to `main`.

## Security & rate limits

- The exporter backs off using `X-RateLimit-Retry-After` and surfaces current `limit` and `remaining` as metrics.
- On `429` responses, it increments `airthings_rate_limited_total` and delays the next poll.

## License

Apache 2.0. See `LICENSE`.
