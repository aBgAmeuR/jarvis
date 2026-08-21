# AGENTS

Homelab ops control plane (Go). CLI client + `jarvisd` on a Unix socket.

## Before coding

Read `docs/context.md` and `docs/requirements.md`. Stay in v0.5 unless asked.

## Rules

- English in repo files.
- Secrets: `provider` + `ref` only. Fingerprint and key path live in the secrets file.
- Reach Proxmox with SSH + `pvesh`/`pct`. Reach Docker with `pct exec` on `docker.vmid`. Do not list guests or containers in YAML.
- No task queue until v1 (updates).
- No auto-retry.
- Small diffs. Match existing layout.

## Docs

Update `docs/` and `README.md` in the same change when behavior or scope shifts.

## Done

Meet `docs/requirements.md` for the current slice. Docs match the binaries.
