
VERSION = $(shell git describe --tags --always)
FLAGS = -ldflags "\
  -X main.VERSION=$(VERSION) \
"

.PHONY: run
run:
	go run $(FLAGS) main.go

.PHONY: build
build:
	CGO_ENABLED=0 go build $(FLAGS) -o bin/seniorgocms main.go

.PHONY: clean
clean:
	rm -f bin/*

.PHONY: deps
deps:
	go get -t -u ./...
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test -cover ./...

.PHONY: version
version:
	@echo $(VERSION)
