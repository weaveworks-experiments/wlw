#!/bin/bash
set -xe
CGO_ENABLED=0 go build -o server server.go
docker build -t lmarsden/wlw:faster .
docker push lmarsden/wlw:faster
