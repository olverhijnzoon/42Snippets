BINARY=golangtime
build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

god: build_go run_go

clean_go:
	rm -f $(BINARY)

CLOCK=clock
build_clock:
	go build -o $(CLOCK)X $(CLOCK).go

run_clock:
	./$(CLOCK)X

god_clock: build_clock run_clock

clean_clock:
	rm -f $(CLOCK)

clean_binaries:
	rm -f ./*X