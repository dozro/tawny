FROM alpine:3.22.2 AS buildenv

RUN apk add go
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o tawnyfm ./cmd/fm-proxy/main.go

FROM scratch

EXPOSE 8080

WORKDIR /app
COPY ./assets /app/assets
COPY ./api /app/api
COPY --from=buildenv /build/tawnyfm /app/tawnyfm

ENTRYPOINT ["/app/tawnyfm"]