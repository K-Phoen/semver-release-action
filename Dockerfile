# Build
FROM golang:1.13-alpine AS build

WORKDIR /go/src/app
ADD ./ /go/src/app

RUN go build -o bumper

# Run
FROM alpine:3.10

COPY --from=build /go/src/app/bumper /bumper
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
