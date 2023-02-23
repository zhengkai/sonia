#!/bin/bash -e

DIR=$(readlink -f "$0") && DIR=$(dirname "$DIR") && cd "$DIR" || exit 1

. ./common.sh
LOG="log file: $LOG_FILE"

PID=$(./get-pid.sh "$TYPE" 2>/dev/null || :)
if [ -n "$PID" ]; then
	echo "server running, pid = $PID"
	echo "$LOG"
	exit
fi

ulimit -n 65535 >/dev/null 2>&1 || :
echo "ulimit = $(ulimit -n)"

nohup "$EXE" > "$LOG_FILE" 2>&1 &
PID="$!"
echo "$PID" > "$PID_FILE"
echo "new server started, pid = $PID"
echo "$LOG"
