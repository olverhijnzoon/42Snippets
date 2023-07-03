SCRIPT_NAME=trace.sh
DOCKER_IMAGE_NAME=trace-image
DOCKER_CONTAINER_NAME=trace-container

all:
	chmod +x $(SCRIPT_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) ls

trace_ls: all docker-build docker-run

.PHONY: all trace_ls docker-build docker-run
