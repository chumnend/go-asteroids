all: build start

.PHONY: build
build:
	@echo "Building..."
	@mkdir -p bin/
	@cd ./bin && go build ../cmd/main.go
	@echo "Build complete."

.PHONY: start
start:
	@echo "Executing..."
	@./bin/main

.PHONY: test
test:
	@echo "Testing cmd/..."
	@go test ./cmd/

.PHONY: clean
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin
	@echo "Clean complete."