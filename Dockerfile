FROM timeline_package
MAINTAINER  vincent
#指定工作目录
WORKDIR /go/src/TimeLine
RUN mkdir /etc/tl
COPY tl /etc/tl

COPY . .
EXPOSE 8080
ENTRYPOINT ["/bin/bash", "build.sh"]