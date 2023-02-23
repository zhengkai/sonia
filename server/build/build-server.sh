#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR"

GIT_ROOT=$(git rev-parse --show-cdup 2>/dev/null)
if [ -n "$GIT_ROOT" ]; then
	TMPDIR="${GIT_ROOT}/tmp"
fi

. ./common.sh

DATE=$(TZ='Asia/Shanghai' date '+%Y-%m-%d %H:%M:%S')
GO_VERSION=$(go version)
GIT_VERSION=$(./git-hash.sh)
if [ "$DOCKER_RUNNING" == "yes" ] && [ -f .git-hash ]; then
	GIT_VERSION=$(cat .git-hash || :)
	HOSTNAME="docker"
fi

LDFLAGS="-X '${BUILD_PACKAGE}.BuildGoVersion=${GO_VERSION}' \
	-X '${BUILD_PACKAGE}.BuildTime=${DATE}' \
	-X '${BUILD_PACKAGE}.BuildType=${TYPE}' \
	-X '${BUILD_PACKAGE}.BuildHost=${HOSTNAME}' \
	-X '${BUILD_PACKAGE}.BuildGit=${GIT_VERSION}'"

cd ../src

if [ -d "vendor" ]; then
	VENDOR="-mod=vendor"
fi

CGO_ENABLED=0 go build $VENDOR \
	-ldflags "$LDFLAGS" \
	-o "$EXE_NEXT" \
	"../build/${TYPE}/"*.go \
	2> >(while read -r line; do echo -e "\e[38;2;255;45;45;48;2;10;10;10m$line\e[0m" >&2; done)
