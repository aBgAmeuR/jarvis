# Homelab infrastructure

Environment Jarvis is built to operate.

## Hardware and hypervisor

* Host: Proxmox VE
* CPU: Intel Core i3-12100 (12th Gen)
* Memory: 16 GB RAM
* Primary storage: 1 TB NVMe
* Secondary storage: ZFS mirror 2× 4 TB HDD
* Proxmox volumes: `local`, `local-lvm`, `hdd`

## VMs and LXC containers (Proxmox)

| ID | Type | Name | Role |
| --- | --- | --- | --- |
| 100 | LXC | pi-hole | DNS / ad blocking |
| 101 | LXC | cockpit | System administration |
| 102 | LXC | docker-host | Docker runtime and Portainer |
| 103 | LXC | wireguard | VPN / remote access |
| 104 | LXC | postgresql | Central relational database |
| 105 | VM (qemu) | remote-dev | Remote development environment |

## Network, security, and access

* Reverse proxy: Nginx Proxy Manager for HTTP/S routing and TLS
* Auth / SSO: Authelia for single sign-on and forward auth
* External exposure: Cloudflare Tunnel — no inbound ports opened
* Secrets / passwords: Vaultwarden

## Docker environment (LXC 102)

| Stack | Containers | Purpose |
| --- | --- | --- |
| authelia | `authelia` | Access security |
| bentopdf | `bentopdf` | Web PDF editing |
| cloudflare-tunnel | `cloudflared` | Cloudflare network connector |
| filebrowser | `filebrowser` | Web file manager |
| harmony | `harmony-api`, `harmony-web` | Spotify stats app |
| homepage | `homepage` | Homelab landing dashboard |
| nginx-proxy-manager | `nginx-proxy-manager-app-1` | Reverse proxy manager |
| vaultwarden | `vaultwarden` | Password manager |
| ai | `ai-openwebui-1` | AI UI |
| - | `portainer` | Docker admin UI |

## Observability

* Visualization: Grafana
* Metrics: Prometheus
* Logs: Loki via `promtail`
* Tracing: Tempo
* Hypervisor metrics: `pve-exporter` (Proxmox)
* Host metrics: `node-exporter`
* Container metrics: `cadvisor`
* DNS metrics: `pihole-exporter`
