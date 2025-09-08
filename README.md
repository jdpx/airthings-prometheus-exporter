# Airthings Prometheus Exporter

A lightweight Prometheus exporter which periodically fetches sensor data from the Airthings for Consumer API and exposes them as Prometheus metrics. Designed to run in a small container and respect strict API rate limits by polling in the background and serving cached metrics.

## Features

- Background poller with cached metrics (no API calls during `/metrics` scrapes)
- Handles API rate limits and backs off using `X-RateLimit-*` headers
- Exposes rate limit state and error counters as metrics
- Configurable units, poll interval, device filtering
- Multi-arch container images published via GoReleaser

## How it works

- A background worker polls the Airthings API on a schedule, respecting rate limits and pagination.
- Results are cached in-memory and exposed on `/metrics` using Prometheus metrics.
- If `ACCOUNT_ID` is not provided, the exporter discovers the single account automatically.

Client code is generated from the API specification at `openapi.yaml` using `openapi-generator`.

## Configuration (environment variables)

- `AIRTHINGS_TOKEN` (required): Bearer token for the Airthings Consumer API
- `ACCOUNT_ID` (optional): Account ID. If unset, the exporter discovers the single account
- `UNIT` (optional): `metric` (default) or `imperial`
- `POLL_INTERVAL` (optional): Poll interval, e.g. `60s` (will dynamically back off when rate-limited)
- `INCLUDE_SERIALS` (optional): Comma-separated serial numbers to include (default: all devices)
- `LISTEN_ADDR` (optional): HTTP listen address (default `:9000`)
- `LOG_LEVEL` (optional): `info` (default) or `debug`

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

With Go installed:

```bash
AIRTHINGS_TOKEN=your_token \
UNIT=metric \
POLL_INTERVAL=60s \
LISTEN_ADDR=":9000" \
INCLUDE_SERIALS="" \
LOG_LEVEL=info \
go run ./cmd/airthings_exporter
```

Then visit `http://localhost:9000/metrics`.

## Docker

Build locally:

```bash
docker build -t ghcr.io/jdpx/airthings_prometheus_exporter:dev .
```

Run:

```bash
docker run --rm -p 9000:9000 \
  -e AIRTHINGS_TOKEN=your_token \
  -e UNIT=metric \
  ghcr.io/jdpx/airthings_prometheus_exporter:dev
```

### docker-compose

A sample compose file is included:

```bash
docker compose up -d
```

Set the following in your shell or a `.env` file used by Compose:

```bash
AIRTHINGS_TOKEN=your_token
UNIT=metric
POLL_INTERVAL=60s
INCLUDE_SERIALS=
LISTEN_ADDR=:9000
LOG_LEVEL=info
```

## Prometheus scrape config

```yaml
- job_name: airthings
  static_configs:
    - targets: ["airthings-exporter:9000"]
```

## Code generation

We use `openapi-generator` to generate the API client from `openapi.yaml`:

```bash
brew install openapi-generator
openapi-generator generate -i openapi.yaml -g go -o internal/airthings --additional-properties=packageName=airthings
```

## Development

- Go version: 1.24
- Linting: `golangci-lint` via CI
- Build locally: `go build ./...`
- Test: `go test ./...`

## Releases & CI

- Multi-arch images are built and published to GHCR via GoReleaser on pushes to `main`.
- GitHub Actions workflow in `.github/workflows/release.yaml` automatically bumps a patch tag and runs GoReleaser.

## Security & rate limits

- The exporter backs off using `X-RateLimit-Retry-After` and surfaces current `limit` and `remaining` as metrics.
- On `429` responses, it increments `airthings_rate_limited_total` and delays the next poll.

## License

Apache 2.0. See `LICENSE`.
