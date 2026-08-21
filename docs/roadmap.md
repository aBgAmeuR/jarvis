# Roadmap

## Code now

v0.2: Unix socket, `jarvis service ls` via per-host SSH and service YAML. Replace it with the v0.5 commands.

## v0.5 (build this)

- [ ] `jarvisd` SSHes to Proxmox with host-key pinning
- [ ] `jarvis proxmox ls|start|stop` (`pct` / `pvesh`; LXC mutate; no stop on vmid 106; no QEMU mutate)
- [ ] `jarvis docker ls|start|stop` via `pct exec 102` (container grain; `STATE` + `HEALTH`)
- [ ] Error if CT 102 is down
- [ ] Example config: `examples/config/`
- [ ] Drop `jarvis service ls` once the commands above work

## After

| Slice | Work |
| --- | --- |
| v0.7 | Update check only |
| v1 | Apply updates; queue for those jobs |
| v1.5 | VM guest agent; QEMU start/stop |
| v3 | HTTP, then UI |

Out of scope: replacing Proxmox, Kubernetes, Portainer as the Docker source of truth.
