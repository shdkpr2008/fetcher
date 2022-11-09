FROM golang:1.19.3-alpine as build

WORKDIR /build
COPY . .
ENV GOARCH=arm64
ENV GOOS=linux
ENV CGO_ENABLED=1

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake gcc
RUN make clean
RUN make configure
RUN make build

FROM golang:1.19.3-alpine

WORKDIR /app
RUN apk add --no-cache dumb-init
RUN update-ca-certificates
COPY --from=build /build/fetcher /tmp
ENTRYPOINT ["dumb-init", "--", "/tmp/fetcher"]
