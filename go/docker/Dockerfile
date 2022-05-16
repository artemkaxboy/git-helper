FROM golang:1.18-alpine as build-backend

# TODO CI ARGS

ADD app /build/app
WORKDIR /build/app

ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

RUN echo go version: `go version`

RUN \
    go build -o app /build/app

FROM scratch

COPY --from=build-backend /build/app/app /app

ENTRYPOINT ["/app"]
