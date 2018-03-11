FROM busybox
LABEL maintainer="CKEVI <admin@purpl3.net>"
COPY aare-exporter /bin/aare-exporter
COPY ca-certificates.crt /etc/ssl/certs/

USER nobody
EXPOSE 3005
ENTRYPOINT [ "/bin/aare-exporter" ]
