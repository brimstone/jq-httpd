FROM scratch

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY app /

ENV PORT=80

EXPOSE 80

ENTRYPOINT ["/app"]
