#!/bin/bash

TARGET="Freya"

if [ "$HOSTNAME" != "$TARGET" ]; then
	>&2 echo only run in server "$TARGET"
	exit 1
fi

sudo docker stop sonia
sudo docker rm sonia
sudo docker rmi sonia

sudo cat /tmp/docker-sonia.tar | sudo docker load

sudo docker run -d --name sonia \
	--env "TANK_MYSQL=sonia:sonia@tcp(172.17.0.1:3306)/sonia" \
	--env "STATIC_DIR=/tmp" \
	--env "OUTPUT_PATH=/output" \
	--mount type=bind,source=/www/sonia/output,target=/output \
	--mount type=bind,source=/www/sonia/log,target=/log \
	--mount type=bind,source=/www/sonia/static,target=/tmp \
	--restart always \
	sonia
