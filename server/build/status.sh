#!/bin/bash

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR" || exit 1

. ./common.sh

PID=$(./get-pid.sh "$TYPE")

if [ -z "$PID" ]; then
	exit
fi

cat "/proc/${PID}/limits"
echo
cat "/proc/${PID}/status"
echo
cat "/proc/${PID}/io"
echo
echo "more in /proc/${PID}"
