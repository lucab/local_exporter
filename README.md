# local\_exporter

[![Build status](https://travis-ci.com/lucab/local_exporter.svg?branch=master)](https://travis-ci.com/lucab/local_exporter)
[![Container image](https://quay.io/repository/lucab/local_exporter/status)](https://quay.io/repository/lucab/local_exporter)


`local-exporter` bridges between Prometheus and instrumented on-host daemons that do not expose a web-server on their own.

It is meant to run as an unpriviliged container with few bind-mounts, and can bridge to multiple local endpoints:
 * plain metrics textfile
 * Unix-domain socket
 * DBus method

Configuration is done through a single TOML file.

## Quickstart

```
go get -u -v github.com/lucab/local_exporter && local_exporter serve --help
```

A TOML configuration sample (with comments) is available under [examples](dist/examples/).

An automatically built `x86_64` container image is available on [quay.io](https://quay.io/repository/lucab/local_exporter) and can be run as:

```
docker run -p 9598:9598/tcp -v "$PWD/dist/examples/config.toml:/etc/local_exporter/config.toml" -v /run:/host/run quay.io/lucab/local_exporter:master local_exporter serve -vv
```

## Demo

[![asciicast](https://asciinema.org/a/256453.svg)](https://asciinema.org/a/256453)
