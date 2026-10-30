# Requirements

What Jarvis v1 and v2 must do. Each function has an ID (S-01, D-01…) that the [API](api.md#endpoints) refers to.

Jarvis is one web UI for the homelab. It answers "is everything fine?" in 5 seconds and "why not?" in 2 clicks.

## Principles

- Jarvis collects nothing. It reads Prometheus, Loki, Tempo, Alertmanager, the Proxmox API and `catalog.json`. Missing data is fixed in homelab-iac, not in Jarvis.
- Jarvis is stateless. No database, no volume to back up.
- Health and warnings are Prometheus rules with a `severity` label (`critical`, `warning`, `info`). Jarvis reads them from Alertmanager and computes no status of its own.
- Read-only, except the Files module (v2). Config changes go through git, Terraform and Ansible.
- Every object has a detail page with its own URL: node, guest, service, probe, alert, network client, VPN peer, file.
- One time range for the whole page (1 h, 24 h, 7 d, 30 d). Range, tab and selected object live in the URL.
- No notifications. Alertmanager has no receiver.

## Users and access

- Every TinyAuth account can open Jarvis, from v1. All modules are visible to every account, network clients and logs included.
- The `admin` group can read the file audit log and empty the trash. Every other right is the same for all accounts.

How Jarvis knows who is logged in: [architecture.md](architecture.md#identity).

## Interface

- UI in English. Dates and numbers formatted once, in shared helpers.
- Light and dark themes. The system setting picks one; a manual toggle is kept in the browser.
- Compact: 12 px base text, controls about 26 px high, tabular figures.
- Desktop first, readable on mobile: overview and "To handle" first, detail pages in one column, tables become lists.
- Data refreshes on its own, metrics and alerts every 10 s. Logs can be followed live. Intervals per data type: [api.md](api.md#polling).
- ⌘K searches every object, files included.
- Each metric or log has a link to Grafana Explore.

## v1

Shipped in this order, each module deployable alone: Supervision, Availability, Network, Maintenance.

### Supervision

| ID | Function |
|---|---|
| S-01 | Overview: global status, node resources, key counters, and the "To handle" list. That list holds every firing alert and warning from all modules, sorted by severity. |
| S-02 | Node `pve`: CPU, RAM, load, storage pools, ZFS health, SMART, temperatures. |
| S-03 | Guest list: status, VMID, IP, CPU, RAM, disk, uptime, health, managed by IaC or not. |
| S-04 | Guest page: resources, uptime, services, logs, metrics, health, ports, plus links to its URL, Grafana and its file in homelab-iac. |
| S-05 | Service catalog: name, guest, type (systemd, container, site), status, URL, dependencies. |
| S-06 | Service page: status, request metrics when the service is instrumented, logs, probes, host guest. |

### Availability and alerts

| ID | Function |
|---|---|
| D-01 | Probes: status, latency, availability over 24 h, 7 d and 30 d, 90-day availability bars. |
| D-02 | Certificates: expiry date per site. A warning fires under 14 days. |
| D-03 | Active alerts: severity, since when, object, link to its page. |
| D-04 | Alert history, also shown on each object's page. |
| D-05 | Silences, shown only. They are created with `amtool`. |

### Network

| ID | Function |
|---|---|
| R-01 | Clients: every device seen on the LAN, with name, IP, MAC, DHCP lease, last seen, queries and blocks. |
| R-02 | Client page: top domains, blocked domains, activity over time. |
| R-03 | DNS: queries per second, block rate, top domains, top blocked, top clients, upstream latency. |
| R-04 | VPN: peers, last handshake, bytes received and sent. |

### Maintenance

| ID | Function |
|---|---|
| M-01 | apt updates per guest, security and regular counted apart. |
| M-03 | Backups per guest: covered by a job or not, date, size and result of the last backup. |
| M-04 | IaC drift: guests in Proxmox but not in the repo, and the reverse; resources that differ from the tfvars. |
| M-05 | ZFS snapshots of the `hdd` pool: age of the latest one. |

M-02 (open Renovate PRs) is out of v1. GitHub already lists them.

## v2: Files

Replaces Filestash on the same dataset (`/hdd/filebrowser`, 325 GB, about 67 000 files). One shared space for every account. Ships after the four v1 modules. How it goes live: [deployment.md](deployment.md#files-cutover).

| ID | Function |
|---|---|
| F-01 | Browse: list or grid, breadcrumb, sort by name, size, date or type. Path and selection live in the URL. Dotfiles are shown. |
| F-02 | Upload: files and folders by drag and drop or picker, with progress. Chunked and resumable after a cut. No size limit. On a name clash: replace, rename or skip. |
| F-03 | Download: one file, or a selection or folder as a zip built on the fly. |
| F-04 | Organise: new folder, rename, move (drag or "Move to…"), copy, multi-select. |
| F-05 | Trash: delete moves items to a hidden trash folder on the dataset, with a `.json` sidecar that keeps the original path. Restore puts the item back there. Items are purged after 30 days. Only `admin` can empty the trash by hand. The trash folder never shows in listings. |
| F-06 | Preview in the side panel: image thumbnails, PDF, audio, MP4 and WebM video, text and code with highlighting, rendered Markdown. Other video formats can only be downloaded; there is no transcoding. The panel also shows size, dates and type. |
| F-07 | Search by name across the whole space, from the module and from ⌘K, filtered by type, date and size. It walks the filesystem live; there is no index. |
| F-08 | Audit log, `admin` only: who, what (upload, delete, rename, move, restore), when, which path. It goes to Loki and is kept 14 days. |
| F-09 | Usage: space used on the dataset and the heaviest folders, linked to the `hdd` pool in Supervision. |
| F-10 | Conflicts: an action on a file that changed since it was displayed is refused, with a prompt to reload. Listings update live when someone else changes a folder. |
| F-11 | Mobile: everything except drag and drop. Upload works from the phone gallery. |

Symlinks are followed only when they point inside the dataset. Every path is checked server side so no request can leave the dataset.

Out of v2: share links, personal folders, per-folder rights, online editing, file versions in the UI, WebDAV or sync clients.

## Out of scope

Actions on guests, a public status page, topology map, change timeline, MCP server, notifications.
