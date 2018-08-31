FROM golang:alpine

RUN mkdir -p /data /config
RUN mkdir -p /data/templates /data/org /data/static
RUN mkdir -p /go/src/github.com/fuzzy/goorgeous-george

COPY . /go/src/github.com/fuzzy/goorgeous-george/
COPY george.yml /config/
COPY run.sh /

RUN go get -v github.com/fuzzy/goorgeous-george

EXPOSE 8080
  
CMD /run.sh
