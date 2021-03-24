#!/bin/sh

go build \
    -ldflags "-extldflags \"-static\" -X main.version=${DRONE_TAG} -X main.revision=${DRONE_COMMIT_SHA}" \
    -o dist/untouched_${GOOS}_${GOARCH} .