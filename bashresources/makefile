# Define variables
DOCKER_CONTAINER_NAME = resource-monitor-container
DOCKER_IMAGE_NAME = resource-monitor-image
SCRIPT_NAME = resource_monitor.sh

# The all target
all:
	@echo "Makefile for Docker container running resource monitoring and quota management scripts"

# The docker-build target
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# The docker-run target for our resource monitoring script
run_monitor: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME)

# The docker-run target with bash shell
run_shell: all docker-build
	docker run -it --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) /bin/bash

# The docker-clean target
docker-clean:
	docker rmi $(DOCKER_IMAGE_NAME)
