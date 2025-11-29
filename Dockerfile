FROM alpine:3.22.2 AS buildenv

RUN apk add go
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o tawnyfm ./cmd/fm-proxy/main.go

FROM scratch

LABEL org.opencontainers.image.source="https://github.com/dozro/tawny"
LABEL org.opencontainers.image.licenses="Apache-2.0"

EXPOSE 8080

WORKDIR /app
COPY ./assets /app/assets
COPY ./api /app/api
COPY --from=buildenv /build/tawnyfm /app/tawnyfm
COPY --from=buildenv /etc/ssl /etc/ssl

HEALTHCHECK --interval=30s --timeout=5s CMD curl -f http://localhost:3030/healthz || exit 1

ENTRYPOINT ["/app/tawnyfm"]