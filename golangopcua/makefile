BINARY=golangopcua

god: build_go run_go

build_go:
	go build -o $(BINARY)X $(BINARY).go

run_go:
	./$(BINARY)X

clean_go:
	rm -f $(BINARY).go

clean_binaries:
	rm -f ./*X

deploy_ignition:
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean_ignition:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml

.PHONY: build_go run_go god clean_go clean_binaries
