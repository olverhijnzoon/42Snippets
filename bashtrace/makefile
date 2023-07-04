# Define variables
DOCKER_CONTAINER_NAME = trace-container
DOCKER_IMAGE_NAME = trace-image
SCRIPT_NAME = trace.sh

# The all target
all:
	@echo "Makefile for Docker container running strace and ltrace scripts"

# The docker-build target
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# The docker-run target for 'ls' command
trace_ls: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) ls

# The docker-run target for 'top' command
trace_top: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) top
