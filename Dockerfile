# Docker image for tiger2es
# To build, run docker build --rm --tag jmarin/tiger2es .
# A container with a shell can be started by running docker run -t -i jmarin/tiger2es

FROM alpine:latest
MAINTAINER Juan Marin Otero <juan.marin.otero@gmail.com>

RUN apk update && apk add git mercurial go && rm -rf /var/cache/apk/* 

RUN mkdir /usr/go
ENV GOPATH /usr/go
ENV PATH $PATH:$GOPATH/bin

# Install Go dependencies

RUN go get github.com/codegangsta/cli
RUN go get code.google.com/p/ftp4go
RUN go get github.com/jonas-p/go-shp
RUN go get github.com/kpawlik/geojson
RUN go get github.com/mattbaird/elastigo

RUN mkdir -p $GOPATH/src/github.com/jmarin/tiger2es
ADD . $GOPATH/src/github.com/jmarin/tiger2es
WORKDIR $GOPATH/src/github.com/jmarin/tiger2es
RUN go install

CMD ["/usr/go/bin/tiger2es"]
