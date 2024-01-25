.PHONY: all
all:
	@echo "Building all services and docker images..."
	@cd counting-service && make all
	@cd dashboard-service && make all

# Build the golang services found in the directories ./counting-service and ./dashboard-service
build-services:
	@echo "Building services..."
	@cd counting-service && make build
	@cd dashboard-service && make build

.PHONY: build
build: build-services

# Clean the golang services found in the directories ./counting-service and ./dashboard-service
clean-services:
	@echo "Cleaning services..."
	@cd counting-service && make clean
	@cd dashboard-service && make clean

.PHONY: clean
clean: clean-services