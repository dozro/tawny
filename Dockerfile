FROM node:18-alpine AS swagger-bundler

WORKDIR /bundle

RUN npm install --ignore-scripts -g swagger-cli@4.0.4

COPY ./api/ ./api/
RUN swagger-cli bundle ./api/openapi/openapi.yaml -o ./bundled.yaml --type yaml

FROM golang:1.25.4-alpine3.22 AS buildenv

WORKDIR /build
COPY .  .
RUN CGO_ENABLED=0 GOOS=linux go build -o tawnyfm ./cmd/fm-proxy/main.go

FROM scratch

LABEL org.opencontainers.image.source="https://github.com/dozro/tawny"
LABEL org.opencontainers.image.licenses="Apache-2.0"

EXPOSE 8080

ENV TAWNY_RELEASE_MODE=true
ENV TAWNY_DEVELOP_MODE=false
ENV TAWNY_RUNNING_IN_DOCKER=true

WORKDIR /app
COPY ./assets /app/assets
COPY --from=buildenv /build/tawnyfm /app/tawnyfm
COPY --from=buildenv /etc/ssl /etc/ssl
COPY --from=swagger-bundler /bundle/bundled.yaml /app/api/openapi-bundled.yaml

HEALTHCHECK --interval=30s --timeout=5s CMD curl -f http://localhost:3030/healthz || exit 1

ENTRYPOINT ["/app/tawnyfm"]