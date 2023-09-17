BINARY=golangproxy

build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

god: build_go run_go

clean_go:
	rm -f $(BINARY).go

clean_binaries:
	rm -f ./*X

.PHONY: build_go run_go god clean_go clean_binaries
