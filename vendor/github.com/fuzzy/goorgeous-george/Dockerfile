FROM golang:alpine

RUN mkdir -p /config
RUN mkdir -p /go/src/github.com/fuzzy/goorgeous-george

COPY . /go/src/github.com/fuzzy/goorgeous-george/
COPY george.yml /config/
COPY run.sh /

RUN apk update
RUN apk add gcc libpthread-stubs util-linux musl-utils musl-dev musl
RUN env GOPATH=/go go get -v github.com/fuzzy/goorgeous-george

EXPOSE 8080
  
CMD /run.sh
