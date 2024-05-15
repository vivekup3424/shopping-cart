rm:
	@rm -rf build
build:
	@go build -o build/shopping_cart cmd/main.go
test:
	@go test -v ./..
run: build
	@./build/shopping_cart