.PHONY: all
all:
	@echo "Building all services and docker images..."
	@cd heartbeat-service && make all
	@cd dashboard-service && make all

# Build the golang services found in the directories ./counting-service and ./dashboard-service
build-services:
	@echo "Building services..."
	@cd heartbeat-service && make build
	@cd dashboard-service && make build

docker-publish-local:
	@echo "Publishing services docker images to local registry..."
	@cd heartbeat-service && make docker-publish-local
	@cd dashboard-service && make docker-publish-local

.PHONY: build
build: build-services

# Clean the golang services found in the directories ./counting-service and ./dashboard-service
clean-services:
	@echo "Cleaning services..."
	@cd heartbeat-service && make clean
	@cd dashboard-service && make clean

.PHONY: clean
clean: clean-services