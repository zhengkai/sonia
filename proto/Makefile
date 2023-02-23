SHELL:=/bin/bash

server:
	protoc --go_out=../server/src *.proto

client:
	../client/node_modules/protobufjs/bin/pbjs -t static-module -w es6 --es6 --force-number -o ../client/src/pb/pb.js *.proto
	../client/node_modules/protobufjs/bin/pbts -o ../client/src/pb/pb.d.ts ../client/src/pb/pb.js
	sed 's#^import.*$$#const $$protobuf = require("protobufjs/minimal");#g' -i ../client/src/pb/pb.js

all: server client
