FROM golang:latest

COPY ./src/ /go/src/app/

WORKDIR /go/src/app/

RUN apt-get update -y\
    && apt-get install build-essential -y\
    && go mod download \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/pi/main.go \
    && cp app /usr/local/bin/app \
    && chmod 755 /usr/local/bin/app \

CMD ["/usr/local/bin/app"]

EXPOSE 8080