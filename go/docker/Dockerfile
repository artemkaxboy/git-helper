FROM golang:1.18-alpine as build-backend

ADD app /build/app
WORKDIR /build/app

ENV GOFLAGS="-mod=vendor"

RUN echo go version: `go version`

RUN \
    go build -o git-helper /build/app

FROM alpine

COPY --from=build-backend /build/app/git-helper /git-helper

ENTRYPOINT ["/git-helper"]
