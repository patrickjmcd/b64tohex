APP?=b64tohex
PROJECT?=github.com/patrickjmcd/b64tohex

RELEASE?=$(shell git tag --sort=committerdate | tail -1)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}

build: clean
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

install: build
	sudo cp ${APP} /usr/local/bin/${APP}