#!/bin/bash
go build -o server server.go
docker build -t lmarsden/wlw:v0.1 .
docker push lmarsden/wlw:v0.1
