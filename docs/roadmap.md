# Roadmap

Sequenced after V1. Order reflects dependency and value, not a hard calendar.

See [context.md](context.md) for the full product vision and V1 definition of done.

## Now — V1

- [x] Daemon + CLI over Unix socket (v0.1)
- [ ] Global task queue (one task at a time), SQLite persistence
- [ ] Declarative service actions (`update`, `health` / `status`)
- [ ] SSH plugin executing declared commands
- [ ] Secrets contract (`provider` + `ref`) with stub
- [ ] Pilot services: `harmony` (Docker), `cockpit` (LXC)

## Next

1. **Proxmox plugin** — operate LXC/VM lifecycle beyond raw SSH commands
2. **Docker / Portainer plugin** — native compose and container ops
3. **Real secrets** — Vault or OpenBao behind the existing provider contract
4. **Network API** — same task model, reachable beyond the local socket
5. **Observability** — OpenTelemetry for tasks, failures, and latency
6. **Optional UI** — human-friendly view of services, tasks, and logs
7. **Worker pool** — safe concurrency (e.g. per service), still no reckless parallelism

## Later / maybe

- Multi-user / SSO
- Daemon high availability
- Broader service catalog beyond the V1 pilots
- Policy / approval gates for destructive actions
- AI-assisted ops (suggestions only; execution stays explicit)

## Non-goals (for now)

Rewriting the homelab stack, replacing Proxmox, or becoming a general-purpose orchestration platform (Kubernetes-class). Jarvis stays a focused control plane for this homelab’s weekly ops.
