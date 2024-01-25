.PHONY: all
all: build docker-build docker-publish-local

.PHONY: build
build: build-heartbeat

# Build the heartbeat service
build-heartbeat:
	@echo "Building heartbeat-service..."
	@go build -o dist/heartbeat-service main.go

docker-build: build
	@echo "Building heartbeat-service docker image..."
	@docker build -t heartbeat-service .

docker-publish-local: docker-build
	@echo "Publishing heartbeat-service docker image to local registry..."
	@docker tag heartbeat-service localhost:5001/heartbeat-service
	@docker push localhost:5001/heartbeat-service

.PHONY: clean
clean:
	@echo "Cleaning heartbeat-service..."
	@rm -rf dist
