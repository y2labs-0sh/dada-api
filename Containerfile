FROM ubuntu:20.04 AS build-base

ENV GO111MODULE on
# ENV GOPROXY "http://mirrors.aliyun.com/goproxy"
ENV GOPATH "$HOME/go"
ENV PATH "$HOME/go/bin:$PATH:/usr/local/go/bin"

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list  \
    && apt update && apt install build-essential wget -y \ 
    && apt-get autoclean  
RUN wget 'https://golang.org/dl/go1.15.3.linux-amd64.tar.gz' -O '/tmp/go.1.15.3.tar.gz' \
    && tar -xzf '/tmp/go.1.15.3.tar.gz' -C /usr/local/ && rm '/tmp/go.1.15.3.tar.gz' 
    
RUN mkdir -p /home/dada-api

WORKDIR "/home/dada-api"

COPY . .

RUN go mod download -x && go generate && rm -rf ./*

VOLUME ["/home/dada-api"]

CMD ["go", "build", "-o", "./build/server_ubuntu"]
