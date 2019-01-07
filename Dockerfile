FROM golang
MAINTAINER  vincent
#指定工作目录
RUN mkdir /etc/tl
ADD /etc/tl /etc/tl

WORKDIR /go/src/TimeLine
COPY . .

CMD ["/bin/bash", "build.sh"]