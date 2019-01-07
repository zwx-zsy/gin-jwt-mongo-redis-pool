FROM golang
MAINTAINER  vincent
#指定工作目录
WORKDIR /go/src/TimeLine
RUN mkdir /etc/tl
COPY tl /etc/tl

COPY . .

CMD ["/bin/bash", "build.sh"]