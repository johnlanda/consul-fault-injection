.PHONY: all
all: build docker-build docker-publish-local

# Bundle web assets
.PHONY: bundle
bundle:
	@echo "Bundling web assets..."
	@rice embed-go

# Build the dashboard service
build-dashboard: bundle
	@echo "Building dashboard-service..."
	@go build -o dist/dashboard-service main.go

.PHONY: build
build: build-dashboard

docker-build: build
	@echo "Building dashboard-service docker image..."
	@docker build -t dashboard-service .

docker-publish-local: docker-build
	@echo "Publishing dashboard-service docker image to local registry..."
	@docker tag dashboard-service localhost:5001/dashboard-service
	@docker push localhost:5001/dashboard-service

# Clean the dashboard service
clean-dashboard:
	@rm -rf dist

# Clean the rice output
clean-rice:
	@rm -rf rice-box.go

.PHONY: clean
clean: clean-dashboard clean-rice
	@echo "Cleaning dashboard-service..."
