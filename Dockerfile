FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /http_test
COPY . .
RUN go env && go build -ldflags="-w -s" -o server .

FROM alpine:latest

LABEL MAINTAINER="wjl@gmail.com"

WORKDIR /http_test

COPY --from=0 /http_test/server ./
# COPY --from=0 /verseland-go/util/comic.ttf ./util/
COPY --from=0 /usr/local/go/lib/time/zoneinfo.zip /opt/zoneinfo.zip

ENV ZONEINFO /opt/zoneinfo.zip

EXPOSE 8877

ENTRYPOINT ./server
