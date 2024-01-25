.PHONY: all
all: build docker-build docker-publish-local

.PHONY: build
build: build-dashboard

# Build the dashboard service
build-dashboard:
	@echo "Building dashboard-service..."
	@go build -o dist/dashboard-service main.go

docker-build:
	@echo "Building dashboard-service docker image..."
	@docker build -t dashboard-service:local .

docker-publish-local: docker-build
	@echo "Publishing dashboard-service docker image to local registry..."
	@docker tag dashboard-service:local localhost:5001/dashboard-service
	@docker push localhost:5001/dashboard-service

.PHONY: clean
clean:
	@echo "Cleaning dashboard-service..."
	@rm -rf dist
