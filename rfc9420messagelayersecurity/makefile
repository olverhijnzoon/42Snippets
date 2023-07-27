BINARY=rfc9420messagelayersecurity
KEYFILE=private.pem

build_go:
	go build -o $(BINARY)X .

run_go:
	./$(BINARY)X

god: build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X

generate_key:
	openssl genrsa -out $(KEYFILE) 2048

test: generate_key
	go test -v
