## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## docs: generate documentation for CLI application commands
.PHONY: docs
docs:
	@go run main.go docs -d ./docs

## readme: generate README.md file
.PHONY: readme
readme: docs
	@cp docs/lch.md README.md
	@echo "## Installation\n\n\`\`\`bash\ngo install github.com/nguyen-duc-loc/leetcode-helper@latest\n\`\`\`" >> README.md
	@echo "README.md has been successfully created"

## test: test all code
.PHONY: test
test:
	@go test -v -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	@go mod tidy
	@go mod verify
	@echo 'Vendoring dependencies...'
	@go mod vendor

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Vetting code..."
	@go vet ./...
	@staticcheck ./...
	@echo "Running tests..."
	@go test -v -race -vet=off ./...

## build: build the CLI application
.PHONY: build
build:
	@echo "Building lch for Linux..."
	@GOOS=linux GOARCH=amd64 go build -o=./bin/linux_amd64/lch
	@echo "Building lch for Windows..."
	@GOOS=linux GOARCH=amd64 go build -o=./bin/windows_amd64/lch.exe
	@echo "Building lch for MacOS..."
	@GOOS=linux GOARCH=amd64 go build -o=./bin/darwin_amd64/lch