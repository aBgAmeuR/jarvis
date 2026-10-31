# API

The rules every endpoint follows, and the list of endpoints. Once `packages/openapi/openapi.yaml` exists, the endpoint tables give way to the Scalar docs published on GitHub Pages, and this file keeps only the rules.

## Contract

`openapi.yaml` is the contract. Every field has a `doc` and an `example` tag. Operations are tagged by module.

## Conventions

All routes are under `/api/v1`. A breaking change goes to `/api/v2`. A removed route first gets `Deprecation` (RFC 9745) and `Sunset` (RFC 8594) headers. `GET /api/v1/meta` returns the version, commit and build date.

Resources are plural nouns in kebab case (`vpn-peers`). Each resource has one canonical URL. Only real children are nested (`/nodes/{node}/disks`); everything else is a filter on its collection (`/services?guest=116`). `restore` is the only verb in the API.

JSON fields and query parameters are camelCase ASCII and start with a letter. Units are part of the name: `memoryUsedBytes`, `uptimeSeconds`. Dates are RFC 3339 in UTC.

Lists return `{ "items": [...] }`. Large lists (logs, clients, files, audit events, alert history) use cursor pagination: `limit` defaults to 100 with a maximum of 1000, and the next page is in a `Link: rel="next"` header. Sort with `sort=-lastSeen`.

Time windows are absolute `start` and `end`. The default is `end` = now and `start` = `end` minus 24 h. The front converts the UI range to these two values, so a shared URL shows the same window.

Every `GET` returns an `ETag` and answers `304` when nothing changed, which keeps polling cheap. A short server cache with `singleflight` sends one Prometheus query for identical concurrent requests.

A resource built from several sources carries `degraded: [{ "source": "loki", "reason": "..." }]` and still answers `200` when one source fails.

Each object has `links` with `webUrl`, `grafanaUrl` and `repoUrl` where they apply.

Errors are `application/problem+json` (RFC 9457).

## Identifiers

| Resource | Id | Example |
|---|---|---|
| Node | Proxmox node name | `pve` |
| Guest | VMID | `116` |
| Service | catalog slug | `grafana` |
| Probe | slug | `http-grafana`, `icmp-116`, `dns-116` |
| Alert | Alertmanager fingerprint | |
| Client | MAC, lowercase, hyphens | `aa-bb-cc-dd-ee-ff` |
| VPN peer | peer name | |
| File | base64url of the path, `root` for the dataset root | |

A file id changes when the file is renamed or moved. `PATCH` returns the new id in `Location`.

## Time series

Series are named per resource and listed in the spec (`series=cpu,memory`). The browser never sends PromQL.

The client sends no point count. The backend picks the Prometheus step: `step = max(15 s, (end - start) / 720)`, rounded to 15 s, 30 s, 1 min, 2 min, 5 min, 15 min, 30 min or 1 h. `rate()` windows are `max(step, 60 s)`, four scrape intervals at least.

| Window | Step | Points |
|---|---|---|
| 1 h | 15 s | 240 |
| 24 h | 2 min | 720 |
| 7 d | 15 min | 672 |
| 30 d | 1 h | 720 |

The response is columnar and does not return the step. A gap is a missing timestamp.

```json
{
  "start": "2026-10-30T10:00:00Z",
  "end": "2026-10-31T10:00:00Z",
  "series": [
    { "name": "cpuUsageRatio", "unit": "ratio", "timestamps": [], "values": [] }
  ]
}
```

## Polling

The front refreshes with TanStack Query `refetchInterval`, paused when the tab is hidden. Metrics and alerts refresh every 10 s, the inventory every 60 s and the catalog every 5 min. Every query of a page shares one `end`, fixed at each refresh. Following logs asks for entries after the last cursor received.

## Endpoints

Paths are relative to `/api/v1`. The Spec column points to [requirements.md](requirements.md).

### Common

| Method | Path | Spec |
|---|---|---|
| GET | `/meta` | version, commit, build date |
| GET | `/me` | user, groups, `isAdmin` |
| GET | `/sources` | status of each data source |
| GET | `/search?q=` | ⌘K, typed hits: guest, service, probe, client, peer, file |

### Supervision

| Method | Path | Spec |
|---|---|---|
| GET | `/overview` | S-01 |
| GET | `/nodes`, `/nodes/{node}` | S-02 |
| GET | `/nodes/{node}/metrics` | S-02 |
| GET | `/nodes/{node}/storages` | S-02 |
| GET | `/nodes/{node}/zpools` | S-02, M-05 (`latestSnapshotAt`) |
| GET | `/nodes/{node}/disks` | S-02, SMART and temperature |
| GET | `/guests`, `/guests/{vmid}` | S-03, S-04, M-01 (apt counts), M-03 (backup coverage) |
| GET | `/guests/{vmid}/metrics` | S-04 |
| GET | `/services`, `/services/{service}` | S-05, S-06 |
| GET | `/services/{service}/metrics` | S-06, rate, errors, latency |
| GET | `/logs` | S-04, S-06; filters `guest`, `service`, `level`, `q` |

The "To handle" list is `GET /alerts?state=firing`, sorted by severity.

### Availability

| Method | Path | Spec |
|---|---|---|
| GET | `/probes`, `/probes/{probe}` | D-01 |
| GET | `/probes/{probe}/metrics` | D-01 |
| GET | `/probes/{probe}/availability?bucket=1d` | D-01, 90-day bars |
| GET | `/certificates`, `/certificates/{site}` | D-02 |
| GET | `/alerts`, `/alerts/{fingerprint}` | D-03, D-04; `state=firing` by default, filters `severity`, `guest`, `service` |
| GET | `/silences`, `/silences/{id}` | D-05 |

Alertmanager does not keep resolved alerts. `state=resolved` reads the `ALERTS` series in Prometheus (90 days) and computes the fingerprint from the labels with Alertmanager's algorithm, so an alert keeps its id after it resolves.

### Network

| Method | Path | Spec |
|---|---|---|
| GET | `/clients`, `/clients/{mac}` | R-01 |
| GET | `/clients/{mac}/domains?status=allowed\|blocked` | R-02 |
| GET | `/clients/{mac}/metrics` | R-02 |
| GET | `/dns` | R-03 |
| GET | `/dns/metrics` | R-03 |
| GET | `/dns/domains?status=allowed\|blocked` | R-03 |
| GET | `/vpn-peers`, `/vpn-peers/{peer}` | R-04 |
| GET | `/vpn-peers/{peer}/metrics` | R-04 |

Top clients are `GET /clients?sort=-queries&limit=10`.

### Maintenance

| Method | Path | Spec |
|---|---|---|
| GET | `/backups?guest=` | M-03 |
| GET | `/backup-jobs` | M-03 |
| GET | `/drift-items` | M-04; `kind` is `missingInIac`, `missingInProxmox` or `resourceMismatch` |

M-01 and M-05 are fields of `/guests` and `/nodes/{node}/zpools`.

### Files (v2)

| Method | Path | Spec |
|---|---|---|
| GET | `/files/{id}` | F-06, metadata and `ETag` |
| GET | `/files/{id}/children` | F-01 |
| GET | `/files?q=&type=&modifiedAfter=&minSize=` | F-07 |
| GET | `/files/{id}/content` | F-03, F-06; supports `Range` |
| GET | `/files/{id}/thumbnail?size=` | F-06 |
| GET | `/files/{id}/archive` | F-03, one folder as a zip |
| POST | `/archives` | F-03, a selection; `201` with `Location: /archives/{token}` |
| GET | `/archives/{token}` | F-03 |
| POST | `/files` | F-04; new folder (`kind: directory`) or copy (`sourceId`), with `onConflict` |
| PATCH | `/files/{id}` | F-04, rename or move; requires `If-Match` |
| DELETE | `/files/{id}` | F-05, moves to trash; requires `If-Match` |
| POST, HEAD, PATCH, DELETE | `/uploads`, `/uploads/{id}` | F-02, tus protocol |
| GET | `/files/{id}/usage` | F-09 |
| GET | `/trash-items`, `/trash-items/{id}` | F-05 |
| POST | `/trash-items/{id}/restore` | F-05 |
| DELETE | `/trash-items/{id}`, `/trash-items` | F-05, admin |
| GET | `/audit-events` | F-08, admin, read from Loki |

The archive token is signed and holds the list of file ids, so nothing is stored. tus keeps partial uploads in a hidden folder on the dataset. The tus handler is mounted outside Huma and documented apart.

F-10 conflicts return `412 Precondition Failed` when the `ETag` no longer matches, `428 Precondition Required` when `If-Match` is missing, and `409 Conflict` when the name is taken and `onConflict` is `fail`.

File content is a stored XSS risk in a shared space. HTML, SVG and XML are served with `Content-Disposition: attachment`, and every content response has `X-Content-Type-Options: nosniff` and `Content-Security-Policy: sandbox`. Serving content from a separate origin is the stronger option. One function maps a file id to a path inside `JARVIS_FILES_ROOT` and rejects anything outside it. It gets a fuzz test before Files ships.
