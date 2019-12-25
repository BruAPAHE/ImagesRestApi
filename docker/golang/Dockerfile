FROM ubuntu:latest

ENV GOVERSION 1.13
ENV GOROOT /opt/go
ENV GOPATH /root/.go

RUN apt-get update  \
    && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
        build-essential \
        wget \
        curl \
        sudo \
        net-tools \
        git-core \
        logrotate \
        ca-certificates \
        unzip \
        mc \
        apt-transport-https \
        software-properties-common \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN cd /opt && wget https://storage.googleapis.com/golang/go${GOVERSION}.linux-amd64.tar.gz && \
    tar zxf go${GOVERSION}.linux-amd64.tar.gz && rm go${GOVERSION}.linux-amd64.tar.gz && \
    ln -s /opt/go/bin/go /usr/bin/ && \
    mkdir $GOPATH

COPY ./src/ /go/src/app/

WORKDIR /go/src/app/

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/core/main.go \
    && cp app /usr/local/bin/app \
    && chmod 755 /usr/local/bin/app

CMD ["tail","-f","/dev/null"]

EXPOSE 8080