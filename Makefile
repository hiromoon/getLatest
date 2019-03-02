.PHONY: build
build:
	go build -o _build/glv main.go

.PHONY: test
test:
	go test ./...
