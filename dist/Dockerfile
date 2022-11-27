# build environment
FROM golang:1.12-alpine3.9 AS build-env
ADD . /src
RUN cd /src && go build -mod=vendor

# run environment
FROM alpine:3.17
WORKDIR /
COPY --from=build-env /src/local_exporter /usr/local/bin/local_exporter
CMD [ "/usr/local/bin/local_exporter", "serve" ]
