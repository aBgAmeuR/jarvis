# syntax=docker/dockerfile:1

FROM docker.io/library/node:24-bookworm-slim AS web
RUN npm install -g pnpm@12.14.0
WORKDIR /src
COPY . .
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile
RUN pnpm --filter "@jarvis/web..." build

FROM docker.io/library/golang:1.26.8-bookworm AS server
ARG VERSION=dev
ARG COMMIT=unknown
WORKDIR /src
COPY go.work go.work.sum ./
COPY apps/server apps/server
COPY --from=web /src/apps/web/dist/client apps/server/internal/platform/webui/static
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    cd apps/server && CGO_ENABLED=0 go build -trimpath -buildvcs=false \
      -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
      -o /out/jarvis ./cmd/jarvis-server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=server /out/jarvis /usr/local/bin/jarvis
EXPOSE 8080
ENTRYPOINT ["jarvis"]
