run:
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml

clean:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml

#kubectl port-forward svc/ignition-service 8088:8088
