

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/seniorgocms main.go

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

