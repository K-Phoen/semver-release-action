# Build
FROM golang:1.13-alpine AS build

WORKDIR /go/src/app
ADD ./bumper /go/src/app

RUN go build -o bumper

# Run
FROM alpine:3.10

RUN apk add --no-cache --update git \
    && rm -f /var/cache/apk/*

COPY --from=build /go/src/app/bumper /bumper
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
