# Build
FROM golang:1.21-alpine AS build

RUN apk add --no-cache --update make \
    && rm -f /var/cache/apk/*

WORKDIR /go/src/app
ADD ./ /go/src/app

RUN make build

# Run
FROM alpine:3.16

COPY entrypoint.sh /entrypoint.sh
COPY --from=build /go/src/app/bumper /bumper

ENTRYPOINT ["/entrypoint.sh"]
