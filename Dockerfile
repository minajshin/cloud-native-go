FROM golang:1.7.4-alpine
MAINTAINER M Shin

ENV SOURCES /go/src/github.com/minashin/cloud-native-go

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go