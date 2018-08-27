FROM golang

RUN mkdir -p /george/{bin,data}
COPY george /george/bin
COPY run.sh /

CMD run.sh
