# drone-template

Inject building environment variables into template files.

[![Build Status](https://travis-ci.org/hhxiao/drone-template.svg?branch=master)](https://travis-ci.org/hhxiao/drone-template)
[![Go Doc](https://godoc.org/github.com/hhxiao/drone-template?status.svg)](http://godoc.org/github.com/hhxiao/drone-template)
[![Go Report](https://goreportcard.com/badge/github.com/hhxiao/drone-template)](https://goreportcard.com/report/github.com/hhxiao/drone-template)
[![MicroBadger](https://images.microbadger.com/badges/image/hhxiao/drone-template.svg)](https://microbadger.com/images/hhxiao/drone-template "Get your own image badge on microbadger.com")

Drone plugin for injecting building environment variables into template files. For the usage information and a
listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build -t hhxiao/drone-template .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-template' not found or does not exist..
```

## Usage

Execute from the working directory:

Injecting building environment variables into template java files

```
docker run --rm \
  -e DRONE_REPO_OWNER=hhxiao \
  -e DRONE_REPO_NAME=drone_template \
  -e DRONE_REPO_LINK=https://github.com/hhxiao/drone_template \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_REF= refs/heads/master \
  -e DRONE_COMMIT_AUTHOR=hhxiao \
  -e DRONE_COMMIT_AUTHOR_EMAIL=hhxiao@gmail.com \
  -e DRONE_COMMIT_MESSAGE="bug fixing" \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://beta.drone.io/hhxiao/drone-template/1 \
  -e DRONE_TAG=1.0.0 \
  -e TEMPLATES=src/main/java/main.java \
  hhxiao/drone-template
```

## Reference
This plugin references a lot from the official **[drone-slack](https://github.com/drone-plugins/drone-slack)** plugin
