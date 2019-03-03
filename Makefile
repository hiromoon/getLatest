.PHONY: build
build:
	go build -o _build/glv main.go

.PHONY: build.linux
build.linux:
	GOOS=linux GOARCH=amd64 go build -o _build/linux/glv main.go && tar -zcvf glv_linux.tar.gz ./_build/linux

.PHONY: build.darwin
build.darwin:
	GOOS=linux GOARCH=amd64 go build -o _build/darwin/glv main.go && tar -zcvf glv_darwin.tar.gz ./_build/darwin

.PHONY: test
test:
	go test ./...
