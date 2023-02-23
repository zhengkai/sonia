#!/bin/bash

DB="sonia"

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

/usr/bin/mysqldump \
	--default-character-set=binary \
	--add-drop-database \
	--add-drop-table \
	--add-locks \
	--hex-blob \
	--quick \
	--skip-dump-date \
	--databases "$DB" \
	> "${DB}-full.sql"
