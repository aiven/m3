FROM alpine:3.11
LABEL maintainer="The M3DB Authors <m3db@googlegroups.com>"

ENV CONFIG=m3aggregator.json

EXPOSE 5000/tcp 6000/tcp 6001/tcp

ADD bin/m3aggregator /bin/m3aggregator

ENTRYPOINT [ "/bin/sh" ]
CMD [ "-c", "/bin/m3aggregator -f /etc/m3aggregator/${CONFIG}" ]
