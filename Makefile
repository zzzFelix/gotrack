.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -v .

.PHONY: install
install:
	go install .