# CLAUDE.md — working in the `airthings-prometheus-exporter` repo

A Prometheus exporter for the **Airthings for Consumer API**. Public repo:
`github.com/jdpx/airthings-prometheus-exporter`.

The `README.md` is the user-facing reference — env vars, auth modes and the
metric list live there. Update it when behaviour changes rather than restating
it here.

## Layout

| Path | What it is |
| --- | --- |
| `cmd/airthings_exporter/main.go` | Entry point; config via `envconfig`, logging via `logrus` |
| `internal/exporter/exporter.go` | Background poller + cached Prometheus metrics |
| `internal/airthings/` | **Generated** API client — do not hand-edit (see below) |
| `openapi.yaml` | The API spec the client is generated from |
| `Dockerfile.runtime` | Used by GoReleaser; copies a prebuilt binary (no in-image compile) |

## Critical rules

1. **⚠️ Every push to `main` cuts a release.**
   `.github/workflows/release.yaml` auto-increments the patch version, pushes
   the tag itself, and publishes to GHCR. So: **never create a tag by hand**
   (you'll collide with CI), and treat any merge to `main` as shipping a
   version. Docs-only changes still burn a version number — that's accepted, not
   a bug to fix.
2. **`internal/airthings/` is generated code.** It's `openapi-generator` 7.15.0
   output (note the `.openapi-generator/`, `docs/`, `model_*.go`, `git_push.sh`
   files). Change `openapi.yaml` and regenerate; don't patch the generated files
   by hand.
3. **`make generate` is stale — don't run it blind.** It invokes `oapi-codegen`
   writing `internal/airthings/client.gen.go`, which is a *different generator*
   producing a different API shape from what's committed, with
   `skip-prune: false`. Running it would leave the package inconsistent. If you
   genuinely need to regenerate, use openapi-generator 7.15.0 to match, or fix
   the Makefile and migrate the callers deliberately as its own change.
4. **Respect the API rate limit.** The whole design exists because the Airthings
   API limit is strict: a background poller caches values so `/metrics` scrapes
   never hit the API. Never make an API call on the scrape path, and keep the
   `X-RateLimit-*` back-off behaviour intact.
5. **Keep label cardinality low.** Labels are deliberately minimal
   (`serial_number`, `sensor_type`, `unit`, `device_name`, `device_type`). Don't
   add per-reading or timestamp-like labels.
6. **Never commit tokens, client IDs or secrets.** Auth is a static
   `AIRTHINGS_TOKEN` or OAuth2 client credentials, both supplied by env only.

## Build / test

```bash
go build ./...
go test ./...
go run ./cmd/airthings_exporter   # needs auth env vars, see README
```

`make test` and `make run` wrap the latter two. CI lints with `golangci-lint`,
and GoReleaser runs `go mod tidy` + `go test ./...` as release pre-hooks — so a
failing test blocks the release, which happens on push to `main`. Get tests
green *before* merging, not after.

## Deployment

Deployed by the homelab repo, not from here: `~/code/jdpx-local`, in
`ansible/host_vars/docker.yml` under `host_docker_config.images` — running on
the `docker` LXC, port `9819:9000`, credentials from vault.

**Naming trap:** the GitHub repo is hyphenated but the published image is
underscored — `ghcr.io/jdpx/airthings_prometheus_exporter`. The image tag there
is **pinned deliberately** (was floating on `:latest`); bump it explicitly in
the homelab repo after a release.
