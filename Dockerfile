FROM scratch
MAINTAINER Eelco Hoekema <ehoekema@gmail.com>
ADD service-discovery service-discovery

EXPOSE 8080
ENTRYPOINT ["/service-discovery"]
