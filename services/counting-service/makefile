.PHONY: all
all: build docker-build docker-publish-local

# Build the counting-service
.PHONY: build
build:
	@echo "Building counting-service..."
	@go build -o dist/counting-service main.go

docker-build: build
	@echo "Building counting-service docker image..."
	@docker build -t counting-service .

docker-publish-local: docker-build
	@echo "Publishing counting-service docker image to local registry..."
	@docker tag counting-service localhost:5001/counting-service
	@docker push localhost:5001/counting-service

# Clean the counting-service
.PHONY: clean
clean:
	@echo "Cleaning counting-service..."
	@rm -rf dist