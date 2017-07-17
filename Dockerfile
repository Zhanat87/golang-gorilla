FROM golang:1.8-alpine

MAINTAINER Iskakov Zhanat <iskakov_zhanat@mail.ru>

ENV SOURCES /go/src/github.com/Zhanat87/golang-gorilla/
COPY . ${SOURCES}

RUN cd ${SOURCES} && \
    apk add --no-cache git && \
    go get -u github.com/kardianos/govendor && govendor sync && \
    apk del git && \
    cp .env.prod .env && \
    go test $(go list ./... | grep -v /vendor) -v -bench=. && \
    CGO_ENABLED=0 go build

WORKDIR ${SOURCES}

CMD ${SOURCES}golang-gorilla

ENV PORT 8080
EXPOSE 8080