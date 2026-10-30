# Configuration

The `JARVIS_*` environment variables, their defaults and how they are read.

Variables are loaded into one Go struct and validated at startup. Defaults are `envDefault` tags on that struct, the only place they are set. A missing or invalid value stops the process with the variable name. A secret can come from a file with the `_FILE` suffix. The effective config is logged at startup with secrets masked.

| Variable | Default | Use |
|---|---|---|
| `JARVIS_LISTEN_ADDR` | `:8080` | listen address |
| `JARVIS_PROMETHEUS_URL` | `http://prometheus:9090` | Prometheus base URL |
| `JARVIS_LOKI_URL` | `http://loki:3100` | Loki base URL |
| `JARVIS_TEMPO_URL` | `http://tempo:3200` | Tempo base URL |
| `JARVIS_ALERTMANAGER_URL` | `http://alertmanager:9093` | Alertmanager base URL |
| `JARVIS_PROXMOX_URL` | required | Proxmox API URL |
| `JARVIS_PROXMOX_TOKEN_FILE` | required | read-only Proxmox API token |
| `JARVIS_CATALOG_PATH` | `/etc/jarvis/catalog.json` | path to `catalog.json` |
| `JARVIS_TRUSTED_PROXY` | required | address allowed to send `Remote-*` headers (`192.168.1.113`) |
| `JARVIS_FILES_ROOT` | none, Files off when unset | dataset root, v2 only |

`JARVIS_TRUSTED_PROXY` has no default because a wrong one would let any client choose its identity.
