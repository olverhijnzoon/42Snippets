BINARY=rayleighscattering
build_go:
	go build -o ./${BINARY}X ./${BINARY}.go

run_go:
	./${BINARY}X

god: build_go run_go

clean_go:
	rm -f ./${BINARY}

build_docker:
	docker build --tag rayleighscattering .

run_docker:
	docker run rayleighscattering

ddocker: build_docker run_docker