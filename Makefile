.PHONY: build
build:
	go build -o _build/glv main.go

.PHONY: build.linux
build.linux:
	GOOS=linux GOARCH=amd64 go build -o _build/linux/glv main.go

.PHONY: build.darwin
build.darwin:
	GOOS=linux GOARCH=amd64 go build -o _build/darwin/glv main.go

.PHONY: test
test:
	go test ./...
