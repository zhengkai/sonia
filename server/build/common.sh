BIN="sonia-server"
BUILD_PACKAGE="project/build"

TYPE="${1:-dev}"

if [ ! -f "${DIR}/${TYPE}/go.mod" ]; then
	echo "no build \"${TYPE}\""
	exit 1
fi

DIST_DIR="$(dirname "$DIR")/dist/${TYPE}"
mkdir -p "$DIST_DIR"

EXE="${DIST_DIR}/${BIN}"
EXE_NEXT="${EXE}-next"

PID_FILE="${EXE}.pid"
LOG_FILE="${DIST_DIR}/server.log"
