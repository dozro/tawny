FROM alpine:latest AS buildenv

RUN apk add go
WORKDIR /build
COPY . .
RUN go build -o tawnyfm ./cmd/fm-proxy/main.go

FROM scratch

EXPOSE 8080

WORKDIR /app
COPY --from=buildenv /build/tawnyfm /app/tawnyfm

ENTRYPOINT ["/app/tawnyfm"]