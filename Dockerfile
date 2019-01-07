FROM golang
MAINTAINER  vincent
#指定工作目录
RUN mkdir /etc/tl
WORKDIR /etc/tl
COPY /etc/tl/ .

WORKDIR /go/src/TimeLine
COPY . .

CMD ["/bin/bash", "build.sh"]