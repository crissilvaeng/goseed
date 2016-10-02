FROM alpine:latest

MAINTAINER Cristina Silva <cristina.silva@openmailbox.org>

WORKDIR "/opt"

ADD .docker_build/goseed /opt/bin/goseed

CMD ["/opt/bin/goseed"]