# docker-buildkit-demo

[![Build Status](https://cloud.drone.io/api/badges/go-training/docker-buildkit-demo/status.svg)](https://cloud.drone.io/go-training/docker-buildkit-demo)

Docker Build is one of the most used features of the Docker Engine. See the [reference](https://docs.docker.com/develop/develop-images/build_enhancements/)

## Prepare

Easiest way from a fresh install of docker is to set the `DOCKER_BUILDKIT=1` environment variable when invoking the docker build command, such as:

```sh
DOCKER_BUILDKIT=1 docker build .
```

set daemon configuration in `/etc/docker/daemon.json` feature to true and restart the daemon:

```json
{
  "debug": true,
  "experimental": true,
  "features": {
    "buildkit": true
  }
}
```

## Build without buildkit

docker build without buildkit. See the Dockerfile

```dockerfile
FROM golang:1.14-alpine

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
COPY foo/foo.go foo/
COPY bar/bar.go bar/

ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o /app -v -tags netgo -ldflags '-w -extldflags "-static"' .

CMD ["/app"]
```

update the source code and run `make build`

```sh
docker build --progress=plain -t appleboy/docker-demo -f Dockerfile .
#14 [10/10] RUN go build -o /app -v -tags netgo -ldflags '-w -extldflags "-s...
#14 0.391 gin/foo
#14 0.403 gin/bar
#14 0.412 github.com/go-playground/locales/currency
#14 0.438 github.com/gin-gonic/gin/internal/bytesconv
#14 0.441 github.com/go-playground/locales
#14 0.449 golang.org/x/sys/unix
#14 0.464 net
#14 0.471 github.com/gin-gonic/gin/internal/json
#14 0.508 github.com/go-playground/universal-translator
#14 0.511 github.com/leodido/go-urn
#14 0.694 github.com/golang/protobuf/proto
#14 0.754 gopkg.in/yaml.v2
#14 1.535 github.com/mattn/go-isatty
#14 1.789 net/textproto
#14 1.790 crypto/x509
#14 1.920 vendor/golang.org/x/net/http/httpproxy
#14 1.978 vendor/golang.org/x/net/http/httpguts
#14 2.019 github.com/go-playground/validator/v10
#14 2.434 crypto/tls
#14 3.043 net/http/httptrace
#14 3.085 net/http
#14 4.211 net/rpc
#14 4.212 github.com/gin-contrib/sse
#14 4.212 net/http/httputil
#14 4.372 github.com/ugorji/go/codec
#14 6.322 github.com/gin-gonic/gin/binding
#14 6.322 github.com/gin-gonic/gin/render
#14 6.517 github.com/gin-gonic/gin
#14 6.819 gin
#14 DONE 7.8s
```

## Build with buildkit

See the dockerfile

```dockerfile
# syntax = docker/dockerfile:experimental
FROM golang:1.14-alpine

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

RUN --mount=type=cache,target=/var/cache/apk apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY main.go .
COPY foo/foo.go foo/
COPY bar/bar.go bar/

ENV GOOS=linux
ENV GOARCH=amd64
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go build -o /app -v -tags netgo -ldflags '-w -extldflags "-static"' .

CMD ["/app"]
```

update the source code and run `make buildkit`

```sh
docker build --progress=plain -t appleboy/docker-buildkit -f Dockerfile.buildkit .
#16 [stage-0 10/10] RUN --mount=type=cache,target=/go/pkg/mod --mount=type=c...
#16 0.381 gin/foo
#16 0.447 gin
#16 DONE 1.2s
```

`7.8s` -> `1.2s` all save `6.6` second
