# syntax=docker/dockerfile:1.4
FROM alpine:3

COPY distributarepo /

ENTRYPOINT ["/distributarepo"]

