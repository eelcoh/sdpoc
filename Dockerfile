FROM scratch
MAINTAINER Eelco Hoekema <ehoekema@gmail.com>
ADD sdpoc sdpoc

EXPOSE 8080
ENTRYPOINT ["/sdpoc"]
