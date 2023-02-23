# builder stage

FROM golang:latest as builder

ARG DOCKER_RUNNING=yes
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -yq protobuf-compiler tzdata ca-certificates
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

COPY . /project

RUN cd /project/proto && make

RUN /project/server/build/build-server.sh prod

# clean stage

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/
COPY --from=builder /project/server/dist/prod/sonia-server-next /sonia-server

RUN apk add --no-cache gzip brotli

ENV TZ=Asia/Shanghai

CMD ["/sonia-server"]
