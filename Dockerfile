FROM alpine:latest

MAINTAINER Patrick Reagan <reaganpr@gmail.com>

WORKDIR "/opt"

ADD .docker_build/api-test /opt/bin/api-test

CMD ["/opt/bin/api-test"]
