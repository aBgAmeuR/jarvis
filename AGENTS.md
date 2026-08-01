# AGENTS

Homelab ops control plane (Go). CLI client + daemon; health is success.

## Before coding

Read `docs/context.md`, `docs/requirements.md`. Prefer V1 scope; do not pull in roadmap 'later' work unless asked.

## Rules

- English only for repo files (code, docs, comments).
- No cleartext secrets — `provider` + `ref` only.
- Core must not hardcode Docker/LXC/Proxmox; declare actions in YAML, execute via plugins.
- One global task at a time in V1; no auto-retry.
- Keep changes small; match existing layout.

## Docs

Update `docs/` + `README.md` in the same change when behavior or scope shifts. No stale docs.

## Done

Meet `docs/requirements.md`. Docs match what shipped.
