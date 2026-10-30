# Architecture

How Jarvis is built and how it runs. Where it runs in the homelab: [deployment.md](deployment.md).

## Stack

| Layer | Choice |
|---|---|
| Monorepo | pnpm workspaces and Turborepo, with Turborepo's native Go support (`futureFlags.experimentalGoWorkspaces`, `go.work` at the root) |
| API | Go, Huma v2 on the standard `net/http` mux (`humago`) |
| API docs | OpenAPI 3.1 generated from the Go types, served by Scalar at `/api/docs` |
| Front | TanStack Start in SPA mode, TanStack Router and TanStack Query |
| API client | Orval: TanStack Query hooks and zod schemas, no mocks |
| UI | `@jarvis/ui`, custom components on Base UI, Tailwind v4 tokens |
| Charts | TanStack Charts only, Canvas renderer for dense series |
| Component docs | Storybook for `@jarvis/ui`, in `apps/storybook` |
| Lint and format | oxlint and oxfmt for TypeScript, `go vet` and `go fmt` for Go |

Dependencies and toolchain: [development.md](development.md).

Not used: TanStack DB, micro-frontends, MSW or any mock, SSE or WebSocket, a database, mise.

## Runtime

Jarvis is one container, `jarvis-api`. Nothing runs on the Proxmox host. The binary serves the API and the SPA, embedded with `go:embed`.

It reaches `prometheus:9090`, `loki:3100`, `tempo:3200` and `alertmanager:9093` by service name, and Proxmox through its HTTPS API with a read-only `PVEAuditor` token.

## Code layout

Each data source (Prometheus, Loki, Tempo, Alertmanager, Proxmox, catalog, filesystem) is a Go interface in `internal/app`, implemented in `internal/source`. Sources return domain types only.

Front modules are `supervision`, `availability`, `network`, `maintenance` and `files`. Each one is lazy-loaded by route and never imports another module.

## Identity

Caddy forwards the `Remote-User` and `Remote-Groups` headers it gets from TinyAuth. A middleware on `/api/` reads them only when the request comes from `JARVIS_TRUSTED_PROXY`. Any other source gets `401`. The `admin` group unlocks the routes marked admin in [api.md](api.md#endpoints), and other accounts get `403` there. Header values are never logged.

## Operations

| Path | Use |
|---|---|
| `/healthz` | the process is up |
| `/readyz` | every source answers |
| `/metrics` | Prometheus metrics of Jarvis itself |

These paths sit at the root, outside the versioned API and outside identity. They are not in `openapi.yaml`.

Logs are JSON through `slog`. Traces go to Tempo with OpenTelemetry. On shutdown the server lets running requests and tus uploads finish.

HTTP responses carry a strict CSP with no external script, `Referrer-Policy` and `X-Content-Type-Options: nosniff`.
