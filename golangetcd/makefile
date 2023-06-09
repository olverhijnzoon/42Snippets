.PHONY: apply-etcd init-etcd delete-etcd describe-etcd status-etcd logs-etcd shell-etcd wait-etcd build-go run-go god clean-go clean-binaries

BINARY=golangetcd

apply-etcd:
	kubectl apply -f etcd.yaml

init-etcd:
	kubectl exec -it etcd-0 -- etcdctl put my-key "Hello, etcd!"
	kubectl exec -it etcd-0 -- etcdctl get my-key
	kubectl exec -it etcd-1 -- etcdctl get my-key
	kubectl exec -it etcd-2 -- etcdctl get my-key

delete-etcd:
	kubectl delete -f etcd.yaml

describe-etcd:
	kubectl describe statefulset etcd

status-etcd:
	kubectl get pods -l app=etcd

logs-etcd:
	kubectl logs -l app=etcd --tail=50

shell-etcd:
	kubectl exec -it etcd-0 -- /bin/sh

wait-etcd:
	kubectl wait --for=condition=Ready pod -l app=etcd --timeout=300s

build-go:
	go build -o $(BINARY)X $(BINARY).go

run-go:
	sleep 10
	kubectl port-forward svc/etcd 2379:2379 & 
	./$(BINARY)X &

change-config:
	sleep 10
	kubectl exec -it etcd-0 -- etcdctl put config '{"setting1": "Changed 42Snippet config!", "setting2": 42}'

etcd-demo: build-go apply-etcd run-go change-config

clean-go:
	rm -f $(BINARY).go

clean-binaries:
	rm -f ./*X