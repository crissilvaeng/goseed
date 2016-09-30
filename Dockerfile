FROM alpine:latest

MAINTAINER Cristina Silva <cristina.silva@openmailbox.org>

WORKDIR "/opt"

ADD .docker_build/goddamned /opt/bin/goddamned

CMD ["/opt/bin/goddamned"]