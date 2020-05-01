# syntax=docker/dockerfile:experimental
FROM --platform=$BUILDPLATFORM golang:1.14.2-alpine3.11 as builder

ARG TARGETOS
ARG TARGETARCH
ARG DATE
ARG COMMIT
ARG VERSION

ENV DATE $DATE
ENV COMMIT $COMMIT
ENV VERSION $VERSION
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

RUN apk add --no-cache make
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  make build-app

FROM --platform=$BUILDPLATFORM alpine:3.11

ARG DATE
ARG COMMIT
ARG VERSION
ARG VCS_URL

LABEL maintainer="Wojteek" \
  org.label-schema.name="DynHost" \
  org.label-schema.url="https://hub.docker.com/r/wojteek/dynhost" \
  org.label-schema.description="It updates the DynHost in the Cloudflare or OVH" \
  org.label-schema.build-date=$DATE \
  org.label-schema.vcs-ref=$COMMIT \
  org.label-schema.version=$VERSION \
  org.label-schema.vcs-url=$VCS_URL \
  org.label-schema.vendor="Wojteek" \
  org.label-schema.schema-version="1.0"
  
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app
COPY --from=builder --chown=app /app/bin/dynhost .
RUN chmod +x dynhost
USER app

ENTRYPOINT [ "./dynhost" ]
