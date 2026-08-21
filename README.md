# Jarvis

CLI and daemon for Proxmox LXC and Docker containers on [this homelab](docs/homelab.md).

**v0.5 (docs):** `proxmox` / `docker` `ls|start|stop`. You SSH to LXC 106 and talk to `jarvisd` on a Unix socket.

**v0.2:** `jarvis service ls` over per-host SSH. Implement v0.5 against [examples/config](examples/config/).

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

Edit `examples/config/secrets.yaml` to set the key path and PVE host-key fingerprint.

## Docs

- [context](docs/context.md)
- [requirements](docs/requirements.md)
- [roadmap](docs/roadmap.md)
- [homelab](docs/homelab.md)

## License

[MIT](LICENSE)
