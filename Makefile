
fmt:
	goimports -w *.go cmd

build-assets:
	go-bindata -pkg=assets -prefix=assets/swagger-ui -o=./assets/assets.go ./assets/swagger-ui/...

build: build-assets
	go build -o build/swagger-doc ./cmd/swagger-doc

server: build
	./build/swagger-doc $(ARG)

release: build
