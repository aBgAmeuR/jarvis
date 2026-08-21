# Homelab

## Node

Proxmox VE, Intel i3-12100, 16 GB RAM, 1 TB NVMe, ZFS mirror 2×4 TB. Volumes: `local`, `local-lvm`, `hdd`. Node address for Jarvis: `192.168.1.40`.

## Guests

| ID | Type | Name | Role |
| --- | --- | --- | --- |
| 100 | LXC | pi-hole | DNS |
| 101 | LXC | cockpit | Host UI |
| 102 | LXC | docker-host | Docker |
| 103 | LXC | wireguard | VPN |
| 104 | LXC | postgresql | Database |
| 105 | VM | remote-dev | Dev VM (list only in v0.5) |
| 106 | LXC | jarvis | `jarvisd` + CLI (unprivileged, static IPv4, Pi-hole name `jarvis`) |

Create 106 before running v0.5. On the node: user `jarvis`, NOPASSWD sudo for `/usr/sbin/pct` and `/usr/bin/pvesh`, SSH key + host-key fingerprint in the secrets file on 106.

## Docker (CT 102)

| Project | Containers |
| --- | --- |
| authelia | `authelia` |
| bentopdf | `bentopdf` |
| cloudflare-tunnel | `cloudflared` |
| filebrowser | `filebrowser` |
| harmony | `harmony-api`, `harmony-web` |
| homepage | `homepage` |
| nginx-proxy-manager | `nginx-proxy-manager-app-1` |
| vaultwarden | `vaultwarden` |
| ai | `ai-openwebui-1` |
| (none) | `portainer` |

Health for `jarvis docker ls` comes from Compose/Docker healthchecks, not from Jarvis YAML.

## Edge

Nginx Proxy Manager, Authelia, Cloudflare Tunnel, Vaultwarden. Metrics stack (Grafana, Prometheus, Loki, Tempo, exporters) is unused by Jarvis until a later slice.
