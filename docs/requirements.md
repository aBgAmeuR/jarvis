# Requirements

What v0.5 must do. Scope: [context.md](context.md). Order: [roadmap.md](roadmap.md).

Operator runs `jarvis` on LXC 106. `jarvisd` owns the Unix socket and the SSH session to Proxmox.

## v0.5

### FR-01 Config

Daemon loads `config.yaml`: Proxmox `address`, `user`, `credentials.provider`, `credentials.ref`; `jarvis.vmid`; `docker.vmid`. Missing fields abort startup with the field name.

### FR-02 Secrets

Stub resolves `ref` to `private_key_path` and `host_key_fingerprint`. Committed YAML has no passwords, API tokens, or private key material.

### FR-03 Host key

SSH to Proxmox checks `host_key_fingerprint`. A mismatch fails the command.

### FR-04 Proxmox list

`jarvis proxmox ls` lists LXC and QEMU guests from the node. Table columns, in order: `ID`, `TYPE`, `NAME`, `POWER`. `TYPE` is `lxc` or `vm` (QEMU). `POWER` is `running` or `stopped`. Missing name prints `-`. Rows are sorted by ID. Templates are included, unmarked. Table only; no flags.

### FR-05 Proxmox start/stop

`jarvis proxmox start` and `stop` run on LXC. A VM (`TYPE` `vm`) returns an error that names the type (`vm`). `stop` on `jarvis.vmid` returns an error and does not call `pct stop`.

### FR-06 Docker list

`jarvis docker ls` lists containers inside `docker.vmid`. Columns: name, `STATE`, `HEALTH`. `HEALTH` is `n/a` when Docker reports no healthcheck.

If that CT is not running, the command fails, names the vmid, and does not start it.

### FR-07 Docker start/stop

`jarvis docker start` and `stop` take a container name or id.

### FR-08 Socket

CLI fails with a clear error if `jarvisd` is down. v0.5 has no HTTP listener.

### FR-09 Auth to the node

Daemon SSHes as `proxmox.user` (`jarvis`). That account uses NOPASSWD sudo for `/usr/sbin/pct` and `/usr/bin/pvesh` (unrestricted subcommands).

### FR-10 Validate

`jarvis validate` asks the daemon to check config and reachability. The CLI prints one line per finished layer (`parse`, `required`, `values`, `secrets`, `fingerprint`, `ssh`, `guests`) as NDJSON events arrive. Layers after a failure are skipped. Exit 0 only if every printed layer is `ok`.

Checks, in order: YAML KnownFields on `config.yaml` and `secrets.yaml`; FR-01 fields; value rules (`provider=stub`, vmids ≥ 100, distinct vmids); secret ref, key path readable; fingerprint `SHA256:` + 43 unpadded base64 (reject example placeholder); SSH with host-key pin; `jarvis.vmid` and `docker.vmid` present on the node (LXC or QEMU).

## Non-functional

| ID | Requirement |
| --- | --- |
| NFR-01 | Remote command timeout: 15s for `ls`, 60s for start/stop |
| NFR-02 | Failures include enough `pct` / `docker` stderr to act without a second SSH |
| NFR-03 | Config and example secrets are readable and fit for git |
| NFR-04 | README and `docs/` describe v0.5. `proxmox start|stop` and `docker` commands are still to build |

## Done (v0.5)

FR-01 through FR-08 and FR-10 hold on the lab in [homelab.md](homelab.md). FR-09 is an ops prerequisite on the Proxmox node (user, sudoers, key). Docs match the commands the binaries expose.

## Later (no ACs yet)

| Slice | Intent |
| --- | --- |
| v0.7 | Report pending CT packages and Docker image drift. No mutation |
| v1 | Apply updates. Persist those jobs in a queue |
| v1.5 | QEMU start/stop via guest agent |
| v3 | HTTP API for the same operations, then UI |
