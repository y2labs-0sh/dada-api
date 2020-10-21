FROM ubuntu

ENV DEBIAN_FRONTEND noninteractive
ENV GO111MODULE on
ENV GOPROXY="http://mirrors.aliyun.com/goproxy/"
ENV GOPATH="$HOME/go"
ENV PATH="$PATH:$GOPATH/bin"

WORKDIR /home/dada-api

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates
RUN apt-get install -y make golang git
RUN go get github.com/securego/gosec/v2/cmd/gosec

COPY ./ ./
RUN make build
