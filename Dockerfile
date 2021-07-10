FROM golang:1.16-alpine AS build
WORKDIR /go/src/app
COPY go.mod *.go sources.json ./
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /build_blocklist

################################################################################

FROM debian:stable-slim
RUN apt-get update -qq \
    && apt-get install -qq --no-install-recommends \
        ca-certificates \
        curl \
        unbound \
        unbound-anchor \
    && rm -rf /var/lib/apt/lists/*

VOLUME /conf

EXPOSE 53/udp
EXPOSE 53/tcp

COPY docker/scripts /scripts
COPY docker/unbound /default/unbound
COPY sources.json /default/blocklist/

COPY --from=build /build_blocklist /usr/bin/

CMD [ "/scripts/start.sh" ]
