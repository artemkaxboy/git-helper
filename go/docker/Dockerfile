FROM golang:1.18-alpine as build-backend

ADD app /build/app
WORKDIR /build/app

ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

RUN echo go version: `go version`

RUN \
    go build -o app /build/app

FROM scratch

ARG VERSION=SNAPSHOT
ARG REVISION=LOCAL
ARG REF_NAME
ARG CREATED

# https://github.com/opencontainers/image-spec/blob/main/annotations.md
LABEL org.opencontainers.image.created=${CREATED}
LABEL org.opencontainers.image.authors="artemkaxboy@gmail.com"
LABEL org.opencontainers.image.url="https://github.com/artemkaxboy/git-helper"
LABEL org.opencontainers.image.documentation="https://github.com/artemkaxboy/git-helper"
LABEL org.opencontainers.image.source="https://github.com/artemkaxboy/git-helper"
LABEL org.opencontainers.image.version=${VERSION}
LABEL org.opencontainers.image.revision=${REVISION}
LABEL org.opencontainers.image.vendor="artemkaxboy@gmail.com"
LABEL org.opencontainers.image.licenses="Apache 2.0"
LABEL org.opencontainers.image.ref.name=${REF_NAME}
LABEL org.opencontainers.image.title="Git Helper"
LABEL org.opencontainers.image.description="Handy tool to find unused branches in your git repository"

COPY --from=build-backend /build/app/app /app

VOLUME /data
WORKDIR /data

ENTRYPOINT ["/app"]
