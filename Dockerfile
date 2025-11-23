FROM alpine:3.0 AS buildenv

RUN apk add go
WORKDIR /build
COPY . .
RUN go build -o tawnyfm ./cmd/fm-proxy/main.go

FROM scratch

EXPOSE 8080

WORKDIR /app
COPY --from=buildenv /build/api /app/api
COPY --from=buildenv /build/tawnyfm /app/tawnyfm

ENTRYPOINT ["/app/tawnyfm"]