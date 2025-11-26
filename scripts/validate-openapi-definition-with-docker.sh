#!/bin/bash

docker run \
  --rm --tty \
  --volume "$PWD:/data:ro" \
  ibmdevxsdk/openapi-validator:latest \
    --ruleset .validator-rules.yaml \
    api/apispec.yaml