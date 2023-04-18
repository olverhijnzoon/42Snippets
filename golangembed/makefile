BINARY=golangembed

compress_image:
	gzip -c 42SnippetsLogo.png > 42SnippetsLogo.png.gz

build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

god: compress_image build_go run_go

clean_go:
	rm -f $(BINARY)

clean_binaries:
	rm -f ./*X
