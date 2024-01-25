.PHONY: all
all: build docker-build docker-publish-local

.PHONY: build
build: build-dashboard

# Build the dashboard service
build-dashboard:
	@echo "Building dashboard-service..."
	@go build -o dist/dashboard-service main.go

docker-build: build
	@echo "Building dashboard-service docker image..."
	@docker build -t dashboard-service .

docker-publish-local: docker-build
	@echo "Publishing dashboard-service docker image to local registry..."
	@docker tag dashboard-service localhost:5001/dashboard-service
	@docker push localhost:5001/dashboard-service

.PHONY: clean
clean:
	@echo "Cleaning dashboard-service..."
	@rm -rf dist
