# build environment
FROM docker.io/golang:1.19-alpine3.16 AS build-env
ADD . /src
RUN cd /src && go get && go build

# run environment
FROM docker.io/alpine:3.16
WORKDIR /
COPY --from=build-env /src/local_exporter /usr/local/bin/local_exporter
CMD [ "/usr/local/bin/local_exporter", "serve" ]
