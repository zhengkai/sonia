SHELL:=/bin/bash

build: git
	sudo docker build -t sonia -f Dockerfile ../..

run: build
	sudo docker run --env "SONIA_MYSQL=sonia:sonia@tcp(172.17.0.1:3306)/sonia" \
		--mount type=bind,source=/www/sonia/tmp,target=/tmp \
		--mount type=bind,source=/www/sonia/server/dist/prod/log,target=/log \
		sonia

install: build
	sudo docker save sonia > docker-sonia.tar
	scp docker-sonia.tar freya:/tmp
	scp install.sh freya:/tmp
	ssh freya "chmod +x /tmp/install.sh && /tmp/install.sh && rm /tmp/install.sh"

git:
	../../server/build/git-hash.sh > ../../server/build/.git-hash

clean:
	sudo docker stop sonia || :
	sudo docker rm sonia || :
	sudo docker rmi sonia || :
	sudo docker image prune --force
