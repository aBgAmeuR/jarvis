# Jarvis

CLI and daemon for Proxmox LXC and Docker containers on [this homelab](docs/homelab.md).

**v0.5:** `docker` `ls|start|stop` still to build. You SSH to LXC 106 and talk to `jarvisd` on a Unix socket.

## Quick start

```bash
# Edit secrets (key path + real PVE host-key fingerprint)
# examples/config/secrets.yaml

go run ./cmd/jarvisd --config-dir examples/config
go run ./cmd/jarvis validate
go run ./cmd/jarvis proxmox ls
go run ./cmd/jarvis proxmox start 104
```

## config

```yaml
proxmox:
  address: 192.168.1.40
  user: jarvis
  credentials:
    provider: stub
    ref: homelab/pve-ssh

jarvis:
  vmid: 106

docker:
  vmid: 102
```

Edit `examples/config/secrets.yaml` to set the key path and PVE host-key fingerprint (`SHA256:` + 43 unpadded base64).

## Docs

- [context](docs/context.md)
- [requirements](docs/requirements.md)
- [roadmap](docs/roadmap.md)
- [homelab](docs/homelab.md)

## License

[MIT](LICENSE)
