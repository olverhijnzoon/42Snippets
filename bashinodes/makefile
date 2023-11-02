# Define variables
DOCKER_CONTAINER_NAME=inode-container
DOCKER_IMAGE_NAME=inode-image
SCRIPT_NAME=inodes.sh

# The 'all' target
all:
	@echo "Makefile for Docker container to display inode information"

# The 'docker-build' target
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# The 'docker-run' target for showing inode information of '/bin' directory
inode_bin: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) /bin

# The 'docker-run' target for showing inode information of '/etc' directory
inode_etc: all docker-build
	docker run --rm --name $(DOCKER_CONTAINER_NAME) $(DOCKER_IMAGE_NAME) ./$(SCRIPT_NAME) /etc
