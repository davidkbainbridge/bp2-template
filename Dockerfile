FROM alpine:3.2
MAINTAINER David Bainbridge <dbainbri.ciena@gmail.com>
ADD bp2-service-docker /root/bp2-service
EXPOSE 8901

ENV NBI_string_port=8901
ENV NBI_string_type=http
ENV NBI_string_publish=true

ADD bp2/hooks /bp2/hooks

RUN ln -s /bp2/hooks/bash /bin/bash \
    && ln -s /bp2/hooks/hook-to-rest /usr/bin/southbound-update \
    && ln -s /bp2/hooks/hook-to-rest /bp2/hooks/heartbeat \
    && ln -s /bp2/hooks/hook-to-rest /bp2/hooks/peer-update \
    && ln -s /bp2/hooks/hook-to-rest /bp2/hooks/peer-status

ENV BP_HOOK_URL_REDIRECT_SOUTHBOUND_UPDATE=http://127.0.0.1:6789/api/v1/hook/southbound-update
ENV BP_HOOK_URL_REDIRECT_PEER_UPDATE=http://127.0.0.1:6789/api/v1/hook/peer-update
ENV BP_HOOK_URL_REDIRECT_PEER_STATUS=http://127.0.0.1:6789/api/v1/hook/peer-status
ENV BP_HOOK_URL_REDIRECT_HEARTBEAT=http://127.0.0.1:6789/api/v1/hook/heartbeat

ENV BP2HOOK_heartbeat=/bp2/hooks/heartbeat
ENV BP2HOOK_peer-update=/bp2/hooks/peer-update
ENV BP2HOOK_peer-status=/bp2/hooks/peer-status

ENTRYPOINT ["/root/bp2-service"]
