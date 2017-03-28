# Gopnik extended with Redis backend
#
FROM dpokidov/gopnik

LABEL maintainer "sarbash.s@gmail.com"

ENV GOPATH=$GOPATH:/gopnik

ADD . /gopnik/src/rediskv
WORKDIR /gopnik

RUN go get github.com/go-redis/redis

RUN sed -i -E 's/\)/\t_ "rediskv"\n\)/' src/plugins_enabled/config.go
RUN gom exec ./build.bash
