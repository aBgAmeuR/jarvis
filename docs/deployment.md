# Deployment

How the image is built and how Jarvis runs in the homelab, with what it needs from homelab-iac.

## Image

`docker/Dockerfile` builds from the repo root: `docker build -f docker/Dockerfile .`. BuildKit reads `docker/Dockerfile.dockerignore`, so there is no root `.dockerignore`.

The build is multi-stage: `turbo prune` and the Vite build, then `go build -trimpath` with the version set by `-ldflags`, then a distroless image running as non-root.

- Base images are pinned by digest in each `FROM`. Renovate bumps them.
- The only `ARG`s are `VERSION` (default `dev`) and `COMMIT` (default `unknown`). CI sets them to the tag and the SHA.
- The Dockerfile sets no `JARVIS_*` variable. Defaults live in the Go config struct, see [configuration.md](configuration.md).
- The image has no shell. `jarvis-api healthcheck` calls `/healthz` and exits 0 or 1.

## Homelab

Jarvis runs in the `docker/monitoring` compose on CT 111, next to Prometheus, Loki, Tempo and Alertmanager. Caddy publishes it behind TinyAuth `forward_auth`. The guest firewall opens its port to Caddy only.

In the compose file the container runs with `read_only: true`, `cap_drop: [ALL]`, `no-new-privileges`, a memory limit and a healthcheck that runs `jarvis-api healthcheck`. The Files dataset is mounted only when the Files module ships.

The homelab-iac compose pins the image by digest, and Renovate bumps it.

The JSON Schema of `catalog.json` is published from this repo. homelab-iac validates the file against it before deploying.

## Prerequisites in homelab-iac

| Module | Needs |
|---|---|
| Supervision | Alloy on every guest and on the host (done), `catalog.json`, `instance` and `vmid` labels everywhere |
| Availability | blackbox probes, alert rules with `severity` (done), Alertmanager (done) |
| Network | DHCP lease metrics (done), Blocky logs, WireGuard metrics |
| Maintenance | apt metrics (done), Proxmox backup jobs (none exist), read-only Proxmox token, drift exported as a metric, snapshot age metric |
| Files | the dataset mounted in CT 111 and in the `jarvis` container as uid 1000, ZFS snapshot jobs, upload timeouts checked in Caddy |

## Files cutover

1. ZFS snapshots of `hdd` run, hourly and daily. Files does not open before that. Restores from snapshots are done by the admin on the host shell.
2. Files is validated on `jarvis.antoinejosset.fr/files`.
