#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR" || exit 1

. ./common.sh

PID=$(./get-pid.sh "$TYPE" 2>/dev/null || :)
if [ -z "$PID" ]; then
	echo server is not running
	exit
fi

echo 'stoping server' >> "$LOG_FILE" 2>&1 &

echo "kill pid $PID $EXE"
sudo kill "$PID"
while [ -e "/proc/${PID}/exe" ];
do
	sleep 1;
done;
