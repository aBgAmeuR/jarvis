# Jarvis

CLI plus daemon to list, start, and stop Proxmox guests and Docker containers on this homelab.

You SSH to LXC 106 (`jarvis`) and run `jarvis`. That process talks to `jarvisd` on a Unix socket. The daemon SSHs to the Proxmox node as user `jarvis` and runs `pvesh` / `pct`. Docker goes through `pct exec` on CT 102.

The Go tree ships `jarvis validate`. `proxmox` / `docker` commands below are still to implement. Legacy: `jarvis service ls`.

## v0.5

| Command | Behavior |
| --- | --- |
| `jarvis validate` | Live checklist (`✓` / `✗`); on failure, `error:` on the next line |
| `jarvis proxmox ls` | LXC and VMs, power state from the node |
| `jarvis proxmox start\|stop <id>` | LXC only. Refuse QEMU. Refuse stop on `jarvis.vmid` (106) |
| `jarvis docker ls` | Containers on CT 102. Columns `STATE` and `HEALTH` |
| `jarvis docker start\|stop <container>` | One container (not a Compose project) |

`HEALTH` is `healthy`, `unhealthy`, or `n/a` (no Docker healthcheck). `STATE` is the engine status (`running`, `exited`, …).

If CT 102 is stopped, docker commands fail and name that CT. They do not start it.

No task queue. Two mutations can overlap. No update check, no apply, no HTTP.

Config: [examples/config](../examples/config/). You declare the Proxmox SSH target, the Jarvis vmid, and the Docker vmid. `pct` and `docker ps` supply the rest.

Secrets use `provider` + `ref`. The stub maps a ref to an SSH private key path and the PVE host-key fingerprint.

## Later

| Slice | Adds |
| --- | --- |
| v0.7 | Update check (CT packages, image digests). No apply |
| v1 | Apply those updates. Queue only for update jobs |
| v1.5 | QEMU guest agent. VM start/stop |
| v3 | HTTP API, then a UI |

Acceptance: [requirements.md](requirements.md). Order: [roadmap.md](roadmap.md). Lab: [homelab.md](homelab.md).
