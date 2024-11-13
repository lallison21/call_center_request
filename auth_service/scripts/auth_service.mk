PROJECT?=github.com/lallison21/call_center_request/auth_service
NAME?=auth_service
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

.PHONY: auth_build

auth_build:
	go build \
    	-ldflags "-w -s \
    	-X ${PROJECT}/version.Version=${VERSION} \
    	-X ${PROJECT}/version.Name=${NAME} \
    	-X ${PROJECT}/version.Commit=${COMMIT} \
    	-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
    	-o bin/auth_service cmd/auth_service/auth_service.go
