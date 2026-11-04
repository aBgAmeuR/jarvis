# Architecture

How Jarvis is built and how it runs. Where it runs in the homelab: [deployment.md](deployment.md).

## Stack

| Layer | Choice |
|---|---|
| Monorepo | pnpm workspaces and Turborepo, with Turborepo's native Go support (`futureFlags.experimentalGoWorkspaces`, `go.work` at the root) |
| API | Go, Huma v2 on the standard `net/http` mux (`humago`) |
| API docs | OpenAPI 3.1 generated from the Go types, rendered by Scalar in `apps/api`, published on GitHub Pages |
| Front | TanStack Start in SPA mode, TanStack Router and TanStack Query |
| API client | Orval: TanStack Query hooks and zod schemas, no mocks |
| UI | `@jarvis/ui`, custom components on Base UI, Tailwind v4 tokens |
| Icons | `@jarvis/icons`, Hugeicons free set, 14 px with a 1.5 stroke |
| Charts | `@jarvis/charts`, presets and theme over TanStack Charts; Canvas renderer for dense series |
| Tables | TanStack Table, rendered by `DataTable` in `@jarvis/ui` |
| Component docs | Storybook for `@jarvis/ui`, in `apps/storybook` |
| Lint and format | oxlint and oxfmt for TypeScript, `go vet` and `go fmt` for Go |

Dependencies and toolchain: [development.md](development.md).

Not used: TanStack DB, micro-frontends, MSW or any mock, SSE or WebSocket, a database, mise.

## Runtime

Jarvis is one container, `jarvis`. Nothing runs on the Proxmox host. The binary serves the API and the SPA, embedded with `go:embed`.

It reaches `prometheus:9090`, `loki:3100`, `tempo:3200` and `alertmanager:9093` by service name, and Proxmox through its HTTPS API with a read-only `PVEAuditor` token.

## Code layout

The Go code is split by layer, then by feature (`supervision`, `availability`, `identity`, `health`):

| Layer | Holds | May import |
|---|---|---|
| `internal/domain` | types and rules: `Severity`, `SourceError`, `Degraded`, and one package per feature | the standard library |
| `internal/app` | use cases, one package per feature; the interfaces a use case needs (`ports.go`) are defined here | `domain` |
| `internal/source` | one package per data source (Proxmox, Alertmanager, catalog…), implementing those interfaces and returning domain types only | `domain`, `platform` |
| `internal/httpapi` | Huma routes and JSON bodies, one package per feature, and the HTTP middleware | `app`, `domain`, `platform` |
| `internal/platform` | infrastructure without business rules: config, logging, telemetry, the HTTP client with its cache, the server | the standard library |
| `internal/bootstrap` | builds every adapter and use case and wires them; the only package that sees them all | everything |

A feature package never imports another feature of the same layer: when `supervision` needs the firing alerts, it defines the small interface it needs and `bootstrap` provides an adapter on `availability`. The `overview` feature works the same way for the alerts, the Proxmox inventory and the probes. No linter enforces these rules; they are checked in review.

Outgoing calls go through one HTTP client per source: a context deadline on every call, a cap on the body size, TLS 1.2 at least, and a short cache that sends one request for identical concurrent reads (`JARVIS_CACHE_TTL`). A source failure is a `SourceError` with a kind (`timeout`, `unreachable`, `bad_status`, `invalid_response`, `unavailable`); clients see the kind, never the URL or the cause, which go to the logs.

The front is split by page, each one lazy-loaded by route. A page never imports another page; what two pages share goes to `entities` or `shared`.

## Identity

Caddy forwards the `Remote-User` and `Remote-Groups` headers it gets from TinyAuth. A middleware on `/api/` reads them only when the request comes from `JARVIS_TRUSTED_PROXY`. Any other source gets `401`. The `admin` group unlocks the routes marked admin in [api.md](api.md#endpoints), and other accounts get `403` there. Header values are never logged.

## Operations

| Path | Use |
|---|---|
| `/healthz` | the process is up |
| `/readyz` | every source answers |
| `/metrics` | Prometheus metrics of Jarvis itself |

These paths sit at the root, outside the versioned API and outside identity. They are not in `openapi.yaml`.

Logs are JSON through `slog`, each record carrying the `request_id` (also returned in `X-Request-Id`) and the `trace_id`. The query string and the `Remote-*` headers are never logged. Traces go to Tempo with OpenTelemetry when `JARVIS_OTLP_ENDPOINT` is set. A panic in a handler is logged and answered with a 500. API reads carry an `ETag` and answer `304` when the caller already has that version. On shutdown the server lets running requests and tus uploads finish.

HTTP responses carry a strict CSP with no external script, `Referrer-Policy` and `X-Content-Type-Options: nosniff`.
