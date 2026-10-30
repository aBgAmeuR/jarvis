# Development

How to work on the repo: layout, commands, shared tooling, dependencies, CI and releases.

## Repository

```
.claude/CLAUDE.md          points to AGENTS.md
.github/                   workflows, renovate.json, CODEOWNERS
apps/api/                  Go module: cmd/jarvis-api, internal/{domain,app,source,httpapi}
apps/api/openapi.yaml      generated, committed
apps/web/                  TanStack Start SPA, modules in src/modules/<name>
docker/                    Dockerfile, Dockerfile.dockerignore
docs/
packages/ui/               @jarvis/ui, its design tokens and its Storybook
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
| `tooling/oxlint` | `@jarvis/oxlint-config`: `base.json`, `react.json` | a two-line `.oxlintrc.json` with `extends` in each package |
| `tooling/oxfmt` | `@jarvis/oxfmt-config` | `oxfmt -c <path>`; oxfmt 0.75 has no `extends` |
| `tooling/tailwind` | `@jarvis/tailwind-config`: `base.css` with `@import "tailwindcss"`, the dark variant and plugins | `@import "@jarvis/tailwind-config/base.css"` |
| `tooling/golangci` | `.golangci.yml` | `golangci-lint run --config` |
| `tooling/spectral` | API style ruleset | `spectral lint --ruleset` |
| `tooling/commitlint` | commitlint config | `--config` |
| `tooling/go` | Go module with `tool` directives: golangci-lint, gofumpt, oasdiff | `go tool <name>` |

Design tokens (`@theme`: colors, sizes, spacing) live in `packages/ui/src/theme.css`, since they belong to the design system. Storybook, Orval and Vite config stay in the one package that uses each of them.

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

Each tool pins its own version.

| Tool | Pinned in | Installed by |
|---|---|---|
| Go | `toolchain` directive in `apps/api/go.mod` | Go downloads it on first use |
| Node | `devEngines.runtime` in the root `package.json`, `onFail: "download"` | pnpm, with version and checksum in the lockfile |
| pnpm | `packageManager` in the root `package.json` | corepack or `pnpm/action-setup` |
| turbo | root `devDependencies`, from the catalog | pnpm |
| Go tools | `tooling/go/go.mod` | `go tool` |

A machine needs any `go` and any `pnpm`. CI reads the same files: `setup-go` with `go-version-file`, `pnpm/action-setup` with `packageManager`.

## Release

`.github/workflows/release.yml` runs semantic-release on every push to `master`. It reads the commits since the last tag, pushes the next `vX.Y.Z` tag and publishes the GitHub release with its notes. CI never commits: the version lives only in tags. There is no config file and no `CHANGELOG.md`; the plugins are passed on the command line (commit-analyzer, release-notes-generator, github).

## CI (planned)

To set up soon on GitHub Actions. Every check is required before merge.

| Area | Checks |
|---|---|
| Go | golangci-lint v2 with gosec, gofumpt, govulncheck |
| Web | oxlint, oxfmt, `tsc --noEmit` strict, knip, size-limit |
| Contract | `gen` leaves no diff, Spectral with the [API rules](api.md#conventions), oasdiff against `master` |
| Architecture | go-arch-lint for the Go layers, dependency-cruiser for front modules |
| UI | Storybook builds |
| Security | CodeQL for Go and TypeScript, gitleaks, dependency review, Trivy on the image, hadolint |
| Workflows | actionlint, zizmor, actions, least-privilege `GITHUB_TOKEN` |
| Release | Conventional Commits with commitlint, semantic-release, multi-arch image on GHCR with SBOM (syft), cosign signature and build provenance |

Tests are out of scope for now, except the fuzz test on file paths.
