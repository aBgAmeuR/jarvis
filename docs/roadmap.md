# Roadmap

## Code now

Unix socket plus `jarvis validate` and `jarvis proxmox ls`. Replace the remaining gaps with the v0.5 `proxmox` start/stop and `docker` commands.

## v0.5 (build this)

- [x] `jarvis validate` (cobra; NDJSON checklist via daemon; config/secrets/SSH/guest ids)
- [x] `jarvisd` SSHes to Proxmox with host-key pinning (used by validate)
- [x] `jarvis proxmox ls` (`pvesh` list; table `ID TYPE NAME POWER`; type `lxc`/`vm`)
- [ ] `jarvis proxmox start|stop` (LXC mutate; no stop on vmid 106; no QEMU mutate)
- [ ] `jarvis docker ls|start|stop` via `pct exec 102` (container grain; `STATE` + `HEALTH`)
- [ ] Error if CT 102 is down
- [x] Example config: `examples/config/`
- [x] Drop `jarvis service ls`

## After

| Slice | Work |
| --- | --- |
| v0.7 | Update check only |
| v1 | Apply updates; queue for those jobs |
| v1.5 | VM guest agent; QEMU start/stop |
| v3 | HTTP, then UI |

Out of scope: replacing Proxmox, Kubernetes, Portainer as the Docker source of truth.
