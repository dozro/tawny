#!/bin/bash
set -euo pipefail

GIT_DESCRIBE_VERSION="$(git describe --tags --always --long --match 'v[0-9]*.[0-9]*.[0-9]*')"
GIT_LATEST_HASH="$(git rev-parse --short HEAD)"

echo "Building version: $GIT_DESCRIBE_VERSION"
echo "Commit Hash: $GIT_LATEST_HASH"

docker build \
  --build-arg IMAGE_VERSION="$GIT_DESCRIBE_VERSION" \
  --build-arg GIT_REVISION="$GIT_LATEST_HASH" \
  -t tawnyfm:"$GIT_LATEST_HASH" \
  -t tawnyfm:latest-dev \
  -t tawnyfm:"$GIT_DESCRIBE_VERSION" \
  .