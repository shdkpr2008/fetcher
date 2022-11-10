VERSION?="0.0.1"
GOOS?=darwin
GOARCH?=arm64
CGO_ENABLED?=1
LDFLAGS?=""

clean:
	rm -rf fetcher

reset:
	rm -rf *.html
	rm -rf database.db

run:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go run .

configure:
	go mod download
	go mod tidy

build:
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED} go build ${LDFLAGS} .

lint:
	golangci-lint run -v --concurrency 2 \
		--disable-all \
		--timeout 10m \
		--enable gofmt \
		--enable gosimple \
		--enable govet \
		--enable typecheck \
		--enable unused \
		--enable ineffassign \
		--enable staticcheck \
		--enable bodyclose

test:
	gotestsum --format=short-verbose
