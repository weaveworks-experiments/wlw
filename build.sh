#!/bin/bash
set -xe
CGO_ENABLED=0 go build -o server server.go
docker build -t lmarsden/wlw:v0.1 .
docker push lmarsden/wlw:v0.1
