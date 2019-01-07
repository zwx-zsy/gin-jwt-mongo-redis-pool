FROM golang
MAINTAINER  vincent
#指定工作目录
WORKDIR /go/src/TimeLine
COPY . .
RUN mkdir /etc/tl
COPY /etc/tl /etc/tl

CMD ["/bin/bash", "build.sh"]