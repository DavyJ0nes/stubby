# Stubby - Claude Context

## What this project does

Stubby is a small HTTP stub server. Routes, responses, status codes, headers, and query parameters are all defined in a YAML config file. It's used to stand in for real services during development or testing.

## Project structure

```
cmd/main.go                  - entrypoint: loads config, wires up server, handles graceful shutdown
internal/router/route.go     - Route struct definition
internal/router/router.go    - gorilla/mux router setup and http.Handler implementation
internal/config/config.go    - YAML config loading via viper
testing/integration_test.go  - integration tests (build tag: integration)
```

## Common tasks

All tasks are managed via mise:

```bash
mise run test             # unit tests
mise run lint             # golangci-lint
mise run build            # static binary (linux/amd64)
mise run integration-test # docker compose based integration tests
mise run image            # build docker image
mise run clean            # remove built binary
```

## Key decisions

- **Distroless final image** (`gcr.io/distroless/static-debian12:nonroot`): no shell, no package manager, runs as non-root by default. Has no healthcheck tooling inside the container.
- **Config path**: the binary defaults to `config.yaml` (relative). When running in Docker always pass `-config /config.yaml` explicitly as the command — do not rely on working directory.
- **Structured logging**: uses `log/slog` with a JSON handler initialised in `main`. The global default is set so all packages pick it up via `slog.Info/Error`.
- **gorilla/mux Walk**: used at startup to log configured routes. `outputRouteInfo` intentionally blank-identifiers the unused `router` and `ancestors` parameters.

## Testing

- Unit tests: `mise run test` — no external deps
- Integration tests: `mise run integration-test` — requires Docker; builds the image, runs stubby + a Go test container via `docker compose up --exit-code-from sut`
- The `sut` container polls `/dev/tcp/stubby/8080` before running tests to avoid races on startup

## Tool versions

Pinned in `mise.toml`. Go version is also the canonical reference for the Dockerfile builder stage — keep them in sync.
