FROM golang:1.19.3-alpine as build

WORKDIR /build
COPY . .
ENV GOARCH=arm64
ENV GOOS=linux
ENV CGO_ENABLED=1
ENV LDFLAGS="-ldflags '-extldflags \"-static\"'"

RUN apk update && apk upgrade && apk add --update alpine-sdk && \
    apk add --no-cache bash git openssh make cmake gcc
RUN make clean
RUN make configure
RUN make build

FROM isholgueras/chrome-headless:latest

WORKDIR /app
USER root
RUN apt update && apt install dumb-init -y
RUN chown root /usr/lib/chromium/chrome-sandbox
RUN chmod 4755 /usr/lib/chromium/chrome-sandbox
RUN update-ca-certificates
COPY --from=build /build/fetcher /tmp
ENTRYPOINT ["dumb-init", "--", "/tmp/fetcher"]
