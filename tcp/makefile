BINARY=tcp
build_go:
	go build -o $(BINARY)serverX $(BINARY)server.go
	go build -o $(BINARY)clientX $(BINARY)client.go

run_go:
	./$(BINARY)X

god: build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X