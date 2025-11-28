# Building Tawny

## The Server

The proxy server is located in [cmd/fm-proxy](./cmd/fm-proxy).

It is advisable to build it using Docker.

The [Dockerfile](./Dockerfile) is provided at the repository root.

To build it using Docker you can simply run:

```bash
docker build -t github.com/dozro/tawny .
```

## The CLI Tools

### The Last-fm CLI Tool

#### Prerequisites 

- Go **1.24** installed

#### Building

The source code for the last-fm CLI Tool can be found in [cmd/fm-cli](./cmd/fm-cli)

To build it simply run:

```bash
go build -o fmcli ./cmd/fm-cli/main.go
```

### The Tawny CLI Tool

#### Prerequisites

- Go **1.24** installed

#### Building

The source code for the Tawny CLI Tool can be found in [cmd/tawny-cli](./cmd/tawny-cli)

Simply run:

```bash
go build -o tawnycli ./cmd/tawny-cli/.
```