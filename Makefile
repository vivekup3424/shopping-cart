rm:
	@rm -rf build
build:
	@go build -o build/shopping_cart cmd/main.go
test:
	@go test -v ./..
run: build
	@go run ./build/shopping_cart