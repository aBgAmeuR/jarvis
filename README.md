# Jarvis

<p>
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://www.shieldcn.dev/github/license/abgameur/jarvis.svg?variant=secondary&amp;size=xs&amp;mode=dark">
        <img alt="License" src="https://www.shieldcn.dev/github/license/abgameur/jarvis.svg?variant=secondary&amp;size=xs&amp;mode=light">
    </picture>
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://www.shieldcn.dev/github/release/abgameur/jarvis.svg?size=xs&amp;mode=dark&amp;variant=secondary">
        <img alt="Release" src="https://www.shieldcn.dev/github/release/abgameur/jarvis.svg?size=xs&amp;mode=light&amp;variant=secondary">
    </picture>
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://www.shieldcn.dev/github/ci/abgameur/jarvis.svg?variant=secondary&amp;size=xs&amp;mode=dark">
        <img alt="CI" src="https://www.shieldcn.dev/github/ci/abgameur/jarvis.svg?variant=secondary&amp;size=xs&amp;mode=light">
    </picture>
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://shieldcn.dev/flag/fr.svg?size=xs&amp;font=geist&amp;mode=dark" />
        <img alt="built in" src="https://shieldcn.dev/flag/fr.svg?size=xs&amp;mode=light&amp;font=geist" />
    </picture>
</p>

One web console for a Proxmox homelab. Jarvis reads the stack you already run (Prometheus, Loki, Tempo, Alertmanager, the Proxmox API) and shows guests, services, probes, alerts, network clients and maintenance in one place. It collects nothing and stores nothing.

## Stack

Go and [Huma](https://huma.rocks) for the API, TanStack Start for the web app, Base UI and TanStack Charts for the interface. One container serves both.

## Deploy

```bash
docker run -d \
  -e JARVIS_PROXMOX_URL=https://pve:8006 \
  -e JARVIS_PROXMOX_TOKEN_FILE=/token \
  -e JARVIS_TRUSTED_PROXY=192.168.1.10 \
  -v ./token:/token:ro -v ./catalog.json:/etc/jarvis/catalog.json:ro \
  ghcr.io/abgameur/jarvis
```

Jarvis has no login. Run it behind an auth proxy that sets `Remote-User`.

## Develop

Requires Go and pnpm. Both download their pinned versions.

```bash
pnpm install
pnpm dev
```

Configuration is read from `JARVIS_*` environment variables. See [configuration](docs/configuration.md).

## Docs

See [docs/](docs/README.md): requirements, architecture, API, configuration, development and deployment.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
