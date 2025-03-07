# Define variables
BINARY_NAME=orderpack
BUILD_DIR=bin
IMAGE_NAME=orderpack

# Run the application locally
run:
	go run cmd/main.go

# Build the application binary
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go

# Clean up built files
clean:
	rm -rf $(BUILD_DIR)

# Format the code
fmt:
	go fmt ./...

# Run tests
test:
	go test ./...

# Build Docker image
docker-build:
	docker build -t $(IMAGE_NAME) .

# Run Docker container
docker-run:
	docker run -p 8080:8080 $(IMAGE_NAME)

# Stop running containers
docker-stop:
	docker stop $(shell docker ps -q --filter ancestor=$(IMAGE_NAME))

# Remove stopped containers
docker-clean:
	docker rm $(shell docker ps -aq --filter ancestor=$(IMAGE_NAME))

# Remove Docker image
docker-rm:
	docker rmi $(IMAGE_NAME)

# View running containers
docker-ps:
	docker ps --filter ancestor=$(IMAGE_NAME)

# Help command
help:
	@echo "Usage:"
	@echo "  make run          - Run the application locally"
	@echo "  make build        - Build the binary"
	@echo "  make clean        - Remove built files"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run linter"
	@echo "  make test         - Run tests"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make docker-stop  - Stop running containers"
	@echo "  make docker-clean - Remove stopped containers"
	@echo "  make docker-rm    - Remove Docker image"
	@echo "  make docker-ps    - View running containers"