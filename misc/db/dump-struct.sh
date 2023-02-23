#!/bin/bash

DB="sonia"

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

/usr/bin/mysqldump \
	--no-data \
	--default-character-set=binary \
	--add-drop-database \
	--add-drop-table \
	--add-locks \
	--hex-blob \
	--quick \
	--skip-dump-date \
	--databases "$DB" \
	| sed 's# AUTO_INCREMENT=[0-9]*##g' \
	> "${DB}-struct.sql"
