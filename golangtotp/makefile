BINARY=golangtotp
REDIS_CONTAINER_NAME=redis_for_golangtotp

build_go:
	go build -o $(BINARY)X $(BINARY).go

run_redis:
	# Pull the Redis image if not present
	docker pull redis:latest
	# Start Redis container
	docker run --name $(REDIS_CONTAINER_NAME) -d -p 6379:6379 redis:latest

stop_redis:
	# Stop Redis container
	docker stop $(REDIS_CONTAINER_NAME)
	# Remove Redis container
	docker rm $(REDIS_CONTAINER_NAME)

run_go: run_redis
	./$(BINARY)X
	make stop_redis

god: build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X

.PHONY: build_go run_redis stop_redis run_go god clean_go clean_binaries
