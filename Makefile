.DEFAULT_GOAL := run

.PHONY: fmt
fmt:
	@echo "Running go fmt"
	go fmt ./...

.PHONY: vet
vet:
	@echo "Running go vet"
	go vet ./...

.PHONY: build
build: fmt vet
	@echo "Building product-service"
	go build -o product-service .

.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean

.PHONY: test
test: fmt vet
	@echo "Running tests..."
	go test -v ./...

.PHONY: run
run: build
	@echo "Running product-api"
	./product-service