# Development

How to work on the repo: layout, commands, shared tooling, dependencies, CI and releases.

## Repository

```
.github/                   workflows, renovate.json, CODEOWNERS
apps/server/               Go module: cmd/jarvis-server, internal/{domain,app,source,httpapi}
apps/api/                  Scalar API docs
packages/openapi/          openapi.yaml, generated, committed
apps/storybook/            Storybook for @jarvis/ui, with the stories
apps/web/                  TanStack Start SPA, src/{routes,app,pages,entities,shared}
docker/                    Dockerfile, Dockerfile.dockerignore, Dockerfile.prebuilt (CI)
docs/
packages/charts/           @jarvis/charts: chart presets and the chart theme over TanStack Charts
packages/icons/            @jarvis/icons: the Icon component and the Hugeicons free set
packages/ui/               @jarvis/ui and its design tokens
packages/api-client/       Orval config and output in src/gen/
tooling/                   shared config packages
.gitignore  AGENTS.md  LICENSE  README.md
go.work  go.work.sum  package.json  pnpm-lock.yaml  pnpm-workspace.yaml  turbo.json
```

The root holds only the files a tool requires there. pnpm reads its settings from `pnpm-workspace.yaml`, so there is no `.npmrc`. Turborepo's Go support needs `go.work` at the root. There is no `.editorconfig`: oxfmt and `go fmt` own formatting.

## Commands

The root `package.json` has five scripts, each one `turbo run <task>`: `gen`, `lint`, `format`, `build` and `dev`. `dev` runs the API and the Vite dev server, which proxies `/api` to the API.

`pnpm gen` writes `openapi.yaml` from the Go types, then runs Orval. Generated files are committed and never edited by hand, so `build` does not depend on `gen`.

The `go` binary must be on the PATH even for JS-only tasks.

## Tooling

Each folder in `tooling/` is a private workspace package with config only, no code. `pnpm-workspace.yaml` lists `apps/*`, `packages/*` and `tooling/*`.

| Folder | Content | Used with |
|---|---|---|
| `tooling/typescript` | `@jarvis/tsconfig`: `base.json`, `vite-app.json`, `react-library.json` | `"extends": "@jarvis/tsconfig/base.json"` |
| `tooling/oxlint` | `@jarvis/oxlint-config`: `base.json`, `react.json` | a `.oxlintrc.json` in each package, `extends: ["./node_modules/@jarvis/oxlint-config/react.json"]`; oxlint resolves paths, not package names |
| `tooling/oxfmt` | `@jarvis/oxfmt-config`: `.oxfmtrc.json` | a `.oxfmtrc.json` symlink to it in each package; oxfmt has no `extends`, and the Zed extension reads only JSON configs |
| `tooling/tailwind` | `@jarvis/tailwind-config`: `base.css` with `@import "tailwindcss"` | `@import "@jarvis/tailwind-config/base.css"` |
| `tooling/golangci` | `.golangci.yml` | `golangci-lint run --config` |
| `tooling/spectral` | API style ruleset | `spectral lint --ruleset` |
| `tooling/commitlint` | commitlint config | `--config` |
| `tooling/go` | Go module with `tool` directives: golangci-lint (gofumpt and oasdiff to come) | `GOWORK=off go tool -modfile=…/tooling/go/go.mod <name>`; it stays out of `go.work` because its dependencies clash with the server's |

Design tokens (`@theme`: colors, sizes, spacing) live in `packages/ui/src/theme.css`, since they belong to the design system. Storybook is an app of its own in `apps/storybook`, which imports `@jarvis/ui` like `apps/web` does. Orval and Vite config stay in the one package that uses each of them.

## Dependencies

Every npm dependency is declared once, in the default catalog of `pnpm-workspace.yaml`. Packages reference it with `"catalog:"`. `catalogMode: strict` makes `pnpm add` fail for a version outside the catalog, so one version of each dependency exists in the repo. Renovate bumps the catalog. Go dependencies live in `go.mod`.

```yaml
packages:
  - apps/*
  - packages/*
  - tooling/*
catalogMode: strict
catalog:
  react: 19.x.y
  "@tanstack/react-query": 5.x.y
```

Internal packages are referenced with `"workspace:*"`.

## Toolchain

Each tool pins its own version, except Node, which only has a minimum.

| Tool | Pinned in | Installed by |
|---|---|---|
| Go | `go` directive in `go.work` | Go downloads it on first use |
| Node | `engines.node` (`>=24`) in the root `package.json` | the machine or `setup-node` |
| pnpm | `packageManager` in the root `package.json` | corepack or `pnpm/action-setup` |
| turbo | root `devDependencies`, from the catalog | pnpm |
| Go tools | `tooling/go/go.mod` | `go tool` |

A machine needs any `go`, any `pnpm` and Node 24 or later. CI reads the same files: `setup-go` with `go-version-file`, `pnpm/action-setup` with `packageManager`.

## Workflows

| File | Runs on | Does |
|---|---|---|
| `ci.yml` | pull request, `workflow_call` | `server` (lint, build, `gen` leaves no diff), `web`, `storybook` and `image` in parallel, then `ci-ok` (`re-actors/alls-green`), the only required check |
| `image.yml` | `workflow_call` | `build` (the binary on the runner, then `Dockerfile.prebuilt`), then `test` and `scan` (Trivy, HIGH and CRITICAL fail) in parallel. Read-only: this is what pull requests run |
| `publish.yml` | `workflow_call` | `publish` pushes the image that `image.yml` tested in the same run to GHCR, then `attest` attaches the SPDX SBOM (Trivy) and the SLSA provenance, both signed with sigstore, and verifies them. Only `release.yml` calls it, so pull requests never hold a write token |
| `release.yml` | push to `master` | `ci` alongside `version` then `image`; `publish` once all three pass, then `release` |
| `docs.yml` | push to `master` on docs paths | builds Storybook and the API docs, deploys to Pages |

## Release

`release.yml` calls `ci.yml` and `publish` needs it, so nothing is released from a red `master`. `version` and `image` run meanwhile. `version` runs semantic-release in dry-run mode (`cycjimmy/semantic-release-action`) and reads `new_release_version` from its outputs; with no releasable commit the rest is skipped. `image` builds that version and `publish` ships it, and only then does `release` run semantic-release for real: the `vX.Y.Z` tag and the GitHub release exist only if the image was published. CI never commits: the version lives only in tags. There is no config file and no `CHANGELOG.md`; semantic-release runs its default plugins, and the npm one publishes nothing because the root `package.json` is private.

Check an image: `gh attestation verify oci://ghcr.io/abgameur/jarvis:<version> --repo aBgAmeuR/jarvis` for the provenance, and add `--predicate-type https://spdx.dev/Document/v2.3` for the SBOM.

## Checks (planned)

To add to `ci.yml`. Every check is required before merge. Already in place: `gen` leaves no diff, Trivy on the image, the image test, SBOM and provenance signed with sigstore.

| Area | Checks |
|---|---|
| Go | golangci-lint v2 with gosec, gofumpt, govulncheck |
| Web | oxlint, oxfmt, `tsc --noEmit` strict, knip, size-limit |
| Contract | `gen` leaves no diff, Spectral with the [API rules](api.md#conventions), oasdiff against `master` |
| Architecture | go-arch-lint for the Go layers, oxlint `no-restricted-imports` for the front layers |
| UI | Storybook builds |
| Security | CodeQL for Go and TypeScript, gitleaks, dependency review, Trivy on the image, hadolint |
| Workflows | actionlint, zizmor, actions, least-privilege `GITHUB_TOKEN` |
| Release | Conventional Commits with commitlint, multi-arch image |

Tests are out of scope for now, except the fuzz test on file paths.
