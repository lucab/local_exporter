# build environment
FROM docker.io/golang:1.19-alpine3.16 AS build-env
RUN apk --no-cache add git \
  && go install github.com/shift/local_exporter@latest

# run environment
FROM docker.io/alpine:3.16
WORKDIR /
COPY --from=build-env /go/bin/local_exporter /usr/local/bin/local_exporter
CMD [ "/usr/local/bin/local_exporter", "serve" ]
