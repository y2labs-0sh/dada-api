FROM ubuntu

ENV DEBIAN_FRONTEND noninteractive

WORKDIR /home/Aggregator-Info
COPY ./ ./

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates
RUN apt-get install -y make golang

ENV GO111MODULE on
ENV GOPROXY="http://mirrors.aliyun.com/goproxy/"

RUN go build
