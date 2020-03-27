.PHONY: build-app clean install-golint lint

APP_NAME=dynhost
DOCKER=docker
GO=go
GO_INPUT=./cmd/$(APP_NAME)/main.go
GO_OUTPUT=./bin/$(APP_NAME)
GOLINT=$(shell $(GO) list -f {{.Target}} golang.org/x/lint/golint)
VERSION?="0.0.0-dev"
DATE?=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
COMMIT?="$(shell git rev-parse --short HEAD)-dev"

build-app: clean
	@echo "`date +'%d.%m.%Y %H:%M:%S'` Building $(GO_INPUT)"
	$(GO) build \
    	-ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" \
    	-o $(GO_OUTPUT) $(GO_INPUT)
	
clean:
	rm -f $(GO_OUTPUT)
	
install-golint:
	$(GO) get -u golang.org/x/lint/golint

lint:
	$(GOLINT) -set_exit_status ./...
