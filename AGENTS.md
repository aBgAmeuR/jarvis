# AGENTS.md

Jarvis is a read-only web console over an existing homelab observability stack. What it must do: [`docs/requirements.md`](docs/requirements.md). Everything else: [`docs/`](docs/README.md).

## Commands

```bash
pnpm gen      # openapi.yaml from Go types, then the Orval client
pnpm lint     # oxlint, go vet
pnpm format   # oxfmt, go fmt
pnpm build
pnpm dev      # API and Vite dev server, /api proxied to the API
pnpm test
```

Done means `gen`, `lint` and `build` pass and the generated files are part of the change. There is no test suite yet. Do not add one unless asked.

## Rules

- Missing data: stop and name the `homelab-iac` change it needs. Do not add collection here.
- Health is an Alertmanager alert and its `severity`. No thresholds in code.
- The front calls the API only through the Orval hooks.
- Refresh with TanStack Query `refetchInterval`. No SSE or WebSocket.
- UI from `@jarvis/ui`, icons from `@jarvis/icons`, charts from `@jarvis/charts`.

## Quality

Go:

- Handlers only map HTTP to `internal/app`. Logic lives in `internal/app`, types in `internal/domain`.
- Sources return domain types. A Prometheus, Loki or Proxmox type never leaves `internal/source`.
- Layers: `httpapi` → `app` → `domain`; `source` implements the interfaces `app` defines; only `internal/bootstrap` sees them all. Each layer has one package per feature, and a feature never imports another feature of its layer. `depguard` enforces it.
- Small files: one noun or one route per file. Huma schema names are global, so every body and output type has a unique name.
- Every outbound call takes the request `context.Context` and its own timeout.
- Dependencies are passed to constructors and wired in `main`. No package-level state, no `init()`.
- Wrap errors with `%w`. Turn them into Huma errors only in `internal/httpapi`.
- No `panic` on a request path.

Web:

- Components render and handle UI events. Fetching and logic live in hooks.
- Server data lives only in TanStack Query. Never copy it into `useState`, and compute derived values instead of storing them.
- Time range, selection and tab live in search params validated by zod.
- Layers of `apps/web/src`, top to bottom: `routes` (with `router.tsx`), `pages`, `entities`, `shared`. `app` (shell, theme) sits beside `pages` and uses only `shared`. A layer imports only from the layers below it; oxlint enforces it.
- A route file binds a URL to a page and mounts nothing else. A page `pages/<page>` holds everything one URL shows and never imports another page. Its `index.ts` exports the component and `search.ts` the zod schema, so the route keeps the page out of the main chunk.
- `entities/<name>` holds what two pages share about one object: data hooks and small blocks, no layout. Create it at the second use, not before. Generic code goes to `shared`.
- Segments: `ui` renders, `model` holds pure logic and URL state. Components call the Orval hooks directly; a hook in `model` exists only when it adds logic. `router.tsx` sets the inventory interval as the default refetch; a query with another rhythm sets its own.
- Each component fetches what it shows. A page puts one `TimeWindowProvider` above its queries, so they share one `end`. Only the open tab is mounted.
- Lists that can pass 200 rows are virtualized.
- No `any`, no `as` on API data, no `!`, no `@ts-ignore`.

## Ask first

New dependencies, CI or Dockerfile changes, breaking API changes, anything that writes outside the Files module.

## Never

- Commit. The user commits.
- Edit generated files: `packages/openapi/openapi.yaml`, `packages/api-client/src/gen/`, `routeTree.gen.ts`.
- Print or log a token, the content of a `*_FILE` secret or a `Remote-*` header value.
- Reach the Proxmox host over SSH, `pct` or `pvesh`. Use the Proxmox API.
- Add a database or a persistent volume.
