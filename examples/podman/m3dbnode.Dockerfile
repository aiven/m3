FROM alpine:3.11
LABEL maintainer="The M3DB Authors <m3db@googlegroups.com>"

ENV GODEBUG madvdontneed=1
ENV CONFIG_FILE=m3dbnode.json

RUN apk add --no-cache curl jq

# Add m3dbnode binary
ADD bin/m3dbnode /bin/m3dbnode

EXPOSE 2379/tcp 2380/tcp 7201/tcp 7203/tcp 9000-9004/tcp

ENV GODEBUG madvdontneed=1

ENTRYPOINT [ "/bin/sh" ]
CMD [ "-c", "/bin/m3dbnode -f /etc/m3dbnode/${CONFIG_FILE}" ]
