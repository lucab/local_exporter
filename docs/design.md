# Local exporter

## Summary

"local-exporter" bridges between Prometheus and instrumented on-host daemons that do not expose their own web-servers. It builds on the experience of node-exporter textfile collector. It is designed to run in a containerized and unprivileged environment. Additionally, it provides a uniform approach for future instrumentation and exposition of host-local services. 

## Proposal

This proposes building a new exporter, called "local-exporter". It runs (via container orchestration) on each node with host-local instrumented services. It exposes multiple metrics endpoint, one for each service.

## Rationale

Linux hosts are usually running multiple host-local services, most of which not exposing a web server. Such services could be [incrementally instrumented][zincati], but it is unclear how to expose/gather their metrics without adding a web service to each one.

Node-exporter provides a textfile collector, which is however not suited for fanning-in multiple services.

Ideas and notes for this exporter came from an [IRC discussion][irc-logs].

[zincati]: https://github.com/coreos/zincati/issues/27
[irc-logs]: https://matrix.to/#/!HaYTjhTxVqshXFkNfu:matrix.org/$1559744867345463AWGdk:matrix.org?via=matrix.org&via=half-shot.uk&via=mx.geekbundle.org

## Node-exporter shortcomings

 1. It was designed to run privileged and non-containerized
 1. Textfile collector takes a single directory (problematic with different daemons running under their own uid/gid)
 1. Multiple sources are coalesced under the same `/metrics` endpoint
 1. Requires non-timestamped metrics
 1. Relies on daemon-side caches and timers for refreshes

## Design sketch

 1. Written in golang and running from a container (daemonset)
 2. Can be unprivileged, only requires bind-mounts for specific `/run` directories
 3. Can run in overlay/pod network, following cluster network policies
 4. Human configuration file (e.g. TOML)
 5. Exposes a single, fixed port with a web service
 6. Main configuration is a map of "selector" -> "directory"
 7. Exposes its own metrics over `/metrics`
 8. Expose bridged metrics (for other host-local services) as `/bridge?selector=foo`
 9. Prometheus can configure/scrape each service separately
 10. Bridged modes:
     1. dbus (for daemons already speaking dbus): like unix-socket mode, but uses a DBus endpoint 
     1. unix-socket (recommended): bridges a Prometheus GET to a unix-socket read. Daemon spits out metric to each new unix-socket connection.
     1. file (legacy): similar to node-exporter, with caching and timers. However, it doesn't coalesce sources and doesn't complain on timestamps.
 11. Daemons can be directly instrumented (the usual Prometheus way) and write their metrics to via one of the methods above

## Implementation sketch

 * Local exporter: https://github.com/lucab/local_exporter
 * Dummy instrumented local service: https://github.com/lucab/instrumented-daemon
 * Metrics over DBus: https://github.com/lucab/prombus
