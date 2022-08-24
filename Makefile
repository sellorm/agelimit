version = 0.0.1
deps_darwin = main.go ctime_darwin.go
deps_darwin = main.go ctime_linux.go

build-all: build-mac-amd64 build-mac-arm64 build-linux-amd64 build-linux-arm64

build-darwin-amd64: $(deps_darwin)
	[ -d ./builds/darwin/amd64 ] || mkdir -p ./builds/darwin/amd64
	GOOS=darwin GOARCH=amd64 go build -o ./builds/darwin/amd64/ -ldflags "-X main.version=$(version)"

build-darwin-arm64: $(deps_darwin)
	[ -d ./builds/darwin/arm64 ] || mkdir -p ./builds/darwin/arm64
	GOOS=darwin GOARCH=arm64 go build -o ./builds/darwin/arm64/ -ldflags "-X main.version=$(version)"

build-linux-amd64: $(deps_linux)
	[ -d ./builds/linux/amd64 ] || mkdir -p ./builds/linux/amd64
	GOOS=linux GOARCH=amd64 go build -o ./builds/linux/amd64/ -ldflags "-X main.version=$(version)"

build-linux-arm64: $(deps_linux)
	[ -d ./builds/linux/arm64 ] || mkdir -p ./builds/linux/arm64
	GOOS=linux GOARCH=arm64 go build -o ./builds/linux/arm64/ -ldflags "-X main.version=$(version)"

# build-linux-arm: $(deps_linux)
#	[ -d ./builds/linux/arm ] || mkdir -p ./builds/linux/arm
#	GOOS=linux GOARCH=arm go build -o ./builds/linux/arm/ -ldflags "-X main.version=$(version)"

go-env: ## Display current go environment
	go env GOOS GOARCH

go-arch: ## Displays supported target architectures
	go tool dist list

version: ## Display the version number
	@echo "agelimit version: $(version)"

.PHONY: build-all build-linux-arm build-linux-arm64 build-linux-amd64 build-darwin-amd64 build-darwin-amd64 version

clean:
	-rm -r builds
