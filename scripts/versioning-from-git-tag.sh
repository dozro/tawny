#!/bin/bash
set -euo pipefail

export GIT_LATEST_TAG="$(git describe --tags --abbrev=0 --match 'v[0-9]*.[0-9]*.[0-9]*')"
export GIT_DESCRIBE_VERSION="$(git describe --tags --always --long --match 'v[0-9]*.[0-9]*.[0-9]*')"
export GIT_LATEST_HASH="$(git rev-parse --short HEAD)"