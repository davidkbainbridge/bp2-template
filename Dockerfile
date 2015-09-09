FROM alpine:3.2
MAINTAINER David Bainbridge <dbainbri.ciena@gmail.com>
ADD bp2-service-alpine /root/bp2-service
EXPOSE 8901

ENV NBI_string_port=8901
ENV NBI_string_type=http
ENV NBI_string_publish=true

ENTRYPOINT ["/root/bp2-service"]
