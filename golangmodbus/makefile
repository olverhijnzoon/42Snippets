CLIENT_BINARY=golangmodbus_client
SERVER_BINARY=golangmodbus_server

all: build_client build_server run_server run_client

build_client:
	go build -o $(CLIENT_BINARY)X $(CLIENT_BINARY).go

build_server:
	go build -o $(SERVER_BINARY)X $(SERVER_BINARY).go

run_server:
	./$(SERVER_BINARY)X &

run_client:
	./$(CLIENT_BINARY)X

clean:
	rm -f $(CLIENT_BINARY).go $(SERVER_BINARY).go

clean_binaries:
	rm -f ./*X

vet:
	go vet ./...

.PHONY: build_client build_server run_server run_client clean clean_binaries
