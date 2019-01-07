FROM golang
MAINTAINER  jackluo
#指定工作目录
WORKDIR /go/src/ActivitApi
COPY . .

CMD ["/bin/bash", "build.sh"]