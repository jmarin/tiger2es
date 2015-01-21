FROM centos:latest
MAINTAINER Juan Marin Otero <juan.marin.otero@gmail.com>

RUN yum -y update; yum clean all

# Install dependencies

RUN yum -y install git mercurial golang; yum clean all

RUN mkdir /usr/go
ENV GOPATH /usr/go
ENV PATH $PATH:$GOPATH/bin

#RUN mkdir $GOPATH/src
#RUN mkdir $GOPATH/bin
#RUN mkdir $GOPATH/pkg


RUN go get github.com/codegangsta/cli
RUN mkdir -p $GOPATH/src/github.com/jmarin/tiger2es
ADD . $GOPATH/src/github.com/jmarin/tiger2es

CMD ["/bin/bash"]
