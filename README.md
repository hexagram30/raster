# raster

[![Build Status][build-badge]][build]
![Project Status][project-status]
[![Go Report Card][report-card-badge]][report-card]
[![Tagged Version][tag-badge]][tag]

[![Project Logo][logo]][logo-large]

*A raster store service for use in hexagram30 projects*

## About

TBD

## Dependencies

* Non-Go deps: protobuf (e.g., `brew install protobuf`)
* Go-based deps: `make deps`
* Redix: `make redix-run` (see the "Infrastructure" section below)

## Service

Build it:

```shell
$ make
```

Run it:

```shell
$ ./bin/rasterd
```

## Client

Run it:

```shell
$ ./bin/rasterc ping
```
```
PONG
```

```shell
$ ./bin/rasterc version | jq .
```
```
{
  "version": "0.1.0-dev",
  "buildDate": "2019-12-30T06:52:51Z",
  "gitCommit": "1fae8c0",
  "gitBranch": "master",
  "gitSummary": "v0.0.0-3-g1fae8c0-dirty"
}
```

## Infrastructure

The Hexgram30 raster project depends upon the on-disk, key-value store Redix,
which uses the memory-mapped file-based BoltDB as its backend. For development
and testing, a Docker image is provided:

* https://hub.docker.com/r/hexagram30/redixdb

This is a scatch-based Docker image, so it only takes up 16.2 MB uncompressed
(7.78 MB compressed).

To run the image locally:

```shell
$ docker run -it \
		-p 7090:7090 -p 6380:6380 \
		-v `pwd`/data/redixdb:/data \
		hexagram30/redixdb:latest \
		-engine boltdb \
		-storage /data \
		-verbose
```

There is a `make` target provided as a convenience that does the same thing, so
you can just run this:

```shell
$ make redix-run
```

<!-- Named page links below: /-->

[logo]: https://raw.githubusercontent.com/hexagram30/raster/master/assets/images/logo.png
[logo-large]: https://raw.githubusercontent.com/hexagram30/raster/master/assets/images/logo-large.png
[build-badge]: https://github.com/hexagram30/raster/workflows/Go/badge.svg
[build]: https://github.com/hexagram30/raster/actions
[report-card-badge]: https://goreportcard.com/badge/hexagram30/raster?v1
[report-card]: https://goreportcard.com/report/hexagram30/raster
[project-status]: https://img.shields.io/badge/project%20status-planning-violet.svg
[tag-badge]: https://img.shields.io/github/tag/hexagram30/raster.svg
[tag]: https://github.com/hexagram30/raster/tags
