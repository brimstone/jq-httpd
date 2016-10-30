FROM scratch

COPY app /

ENV PORT=80

EXPOSE 80

ENTRYPOINT ["/app"]
