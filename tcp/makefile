BINARY=tcp
build_go:
	go build -o $(BINARY)serverX $(BINARY)server.go
	go build -o $(BINARY)clientX $(BINARY)client.go

run_go:
	chmod +x ./$(BINARY)serverX
	./$(BINARY)serverX & ./$(BINARY)clientX
	lsof -i :8080

god: build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X