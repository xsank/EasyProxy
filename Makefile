build: easyproxy
easyproxy:
	mkdir -p bin
	go build -o 'bin/server' ./src/main.go

dep:
	godep save ./...

vet:
	go tool vet ./

lint:
	golint ./...

test:
	go test -v ./...

.PHONY: build easyproxy dep vet lint test
	
