#!/bin/bash

DB="sonia"

DIR="$(dirname "$(readlink -f "$0")")" && cd "$DIR" || exit 1

mkdir -p table

for T in $(echo "SHOW TABLES" | mysql "$DB" --skip-column-names); do

	/usr/bin/mysqldump \
		--default-character-set=binary \
		--add-drop-database \
		--add-drop-table \
		--add-locks \
		--hex-blob \
		--quick \
		--skip-dump-date \
		"${DB}" "${T}" > "table/${T}.sql"
done
