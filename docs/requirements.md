# Requirements

Functional requirements and acceptance criteria for Jarvis.

- [context.md](context.md) - why / vision / architecture
- [roadmap.md](roadmap.md) - sequencing
- This file - **what must be true** to call a slice done

IDs: `FR-xx` = functional requirement · `AC-xx` = acceptance criterion · `NFR-xx` = non-functional.

Scope labels: **V1** | **Later**.

## Actors

| Actor | Role |
| --- | --- |
| Operator | Runs CLI commands against the local daemon |
| Daemon | Queues, executes, records tasks |
| Plugin | Reaches a target host (V1: SSH) |
| Config author | Declares services, hosts, credentials, actions |

## V1 - Functional requirements

### FR-01 - Declare a service

An operator can describe a service in YAML: name, host, and actions (`update`, `health` at minimum).

**Acceptance**

- [ ] AC-01.1 Service config loads without embedding target-specific logic in the core
- [ ] AC-01.2 Missing required fields fail fast with a clear error
- [ ] AC-01.3 Pilot services `harmony` and `cockpit` are expressible with the same schema

### FR-02 - Update a service

An operator can request `update` for a named service; Jarvis enqueues a task, runs the declared update action, then the declared health check.

**Acceptance**

- [ ] AC-02.1 `jarvis service update <name>` creates a persisted task
- [ ] AC-02.2 Task runs the declared `update` command on the service host
- [ ] AC-02.3 After update, Jarvis runs the declared `health` command
- [ ] AC-02.4 Task ends **success** only if health exits 0
- [ ] AC-02.5 Task ends **failure** if update or health fails; outcome is visible in the CLI
- [ ] AC-02.6 Works end-to-end for `harmony` and `cockpit`

### FR-03 - Status / health without mutation

An operator can check service health without changing the service.

**Acceptance**

- [ ] AC-03.1 `jarvis service status <name>` runs only the declared health action
- [ ] AC-03.2 No update/restart side effects occur
- [ ] AC-03.3 Result is reported as healthy / unhealthy based on health exit code

### FR-04 - Task visibility

Every action is a trackable task with status, logs, and history.

**Acceptance**

- [ ] AC-04.1 Tasks are stored in SQLite
- [ ] AC-04.2 Operator can inspect task status and logs from the CLI
- [ ] AC-04.3 Failed and successful runs remain in history after completion

### FR-05 - Safe concurrency (V1)

Only one task runs at a time (global queue).

**Acceptance**

- [ ] AC-05.1 A second update while one is running is queued, not parallelized
- [ ] AC-05.2 No automatic retry; the operator re-runs manually on failure

### FR-06 - Secrets contract

Credentials are referenced, never stored in cleartext in config files.

**Acceptance**

- [ ] AC-06.1 Config uses `provider` + `ref` (not raw passwords/keys)
- [ ] AC-06.2 V1 ships a stub provider sufficient for local/dev and the pilots
- [ ] AC-06.3 Cleartext secrets are not required in committed YAML

### FR-07 - Plugin boundary

The core does not hardcode Docker or LXC behavior; V1 reaches hosts through an SSH plugin executing declared commands.

**Acceptance**

- [ ] AC-07.1 Same core path serves Docker-hosted and LXC-hosted pilots
- [ ] AC-07.2 Target commands live in service YAML, not in core business logic

### FR-08 - Local control plane access

CLI talks to the daemon over a Unix socket on the Jarvis host.

**Acceptance**

- [ ] AC-08.1 CLI commands fail clearly if the daemon is down
- [ ] AC-08.2 No network API is required for V1

## V1 - Non-functional requirements

| ID | Requirement | Acceptance |
| --- | --- | --- |
| NFR-01 | Tasks have an explicit timeout | Hung remote commands do not block the queue forever |
| NFR-02 | Failures are actionable | CLI/task logs show enough to diagnose without re-SSH guessing |
| NFR-03 | Config is reviewable | Service YAML is human-readable and fit for git |
| NFR-04 | Docs match behavior | README + docs describe the real V1 commands and limits |

## V1 - Definition of done (release gate)

All of the following are true:

1. FR-02 and FR-03 pass on `harmony` and `cockpit`
2. Tasks and logs are inspectable via CLI + SQLite (FR-04)
3. Secrets use the stub contract (FR-06)
4. Out-of-scope items below are not required

### Explicitly out of V1

Network API, UI, Proxmox / Portainer / native Docker plugins, OpenTelemetry, multi-user / SSO, daemon HA, worker pool, real secret backends, auto-retry.

## Later - Requirements (sketch)

Detailed ACs when the roadmap item starts. Intent only:

| ID | Theme | Intent |
| --- | --- | --- |
| FR-20 | Proxmox plugin | Lifecycle ops beyond raw SSH command strings |
| FR-21 | Docker / Portainer plugin | Native compose/container operations |
| FR-22 | Real secrets | Vault / OpenBao behind the same `provider` + `ref` contract |
| FR-23 | Network API | Same task model, reachable remotely |
| FR-24 | Observability | OpenTelemetry for tasks, failures, latency |
| FR-25 | UI | Optional human view of services, tasks, logs |
| FR-26 | Worker pool | Safe concurrency (e.g. per service) |
