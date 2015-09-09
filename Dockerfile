FROM alpine:3.2
MAINTAINER David Bainbridge <dbainbri.ciena@gmail.com>
ADD bp2-service-alpine /root/bp2-service
ENTRYPOINT ["/root/bp2-service"]
