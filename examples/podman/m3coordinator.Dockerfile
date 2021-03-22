FROM alpine:3.11
LABEL maintainer="The M3DB Authors <m3db@googlegroups.com>"

EXPOSE 7201/tcp 7203/tcp
ENV CONFIG_FILE=m3coordinator.json

ADD bin/m3coordinator /bin/m3coordinator

ENTRYPOINT [ "/bin/sh" ]
CMD [ "-c", "/bin/m3coordinator -f /etc/m3coordinator/${CONFIG_FILE}" ]
