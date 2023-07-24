.PHONY: install_argocd port_forward_argocd

# in this snippet, we want to explore ArgoCD
install_argocd:
	kubectl create namespace argocd || true
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# retrieving the admin password and stuff: https://argo-cd.readthedocs.io/en/stable/getting_started/
port_forward_argocd:
	kubectl port-forward svc/argocd-server -n argocd 8080:443

# lets setup a local Gogs with a Postgres
deploy_gogs:
	kubectl apply -f deployment-git.yaml

# be aware of POSTGRES_USER and password from deployment-git.yaml then doing initial Gogs setup
port_forward_gogs:
	kubectl port-forward svc/gogs 3000:3000

# create a user e.g. "test" and get a token to create a repo in local Gogs
create_repo:
	curl -H "Authorization: token 2739647066fec520a075b88430b6398ee08a966b" -H "Content-Type: application/json" -d '{"name": "42snippets_inception"}' http://localhost:3000/api/v1/user/repos

# lets do changes to the copy_deployment-42nginx and push it to 42snippets_inception repo
create_commit:
	# Check if the repository is cloned. If not, clone it.
	if [ ! -d "42snippets_inception" ]; then \
		git clone http://localhost:3000/test/42snippets_inception; \
	fi
	cp copy_deployment-42nginx.yaml 42snippets_inception/deployment-42nginx.yaml
	( \
		cd 42snippets_inception; \
		git add deployment-42nginx.yaml; \
		git commit -m "Update index.html"; \
		git remote add origin http://localhost:3000/test/42snippets_inception.git; \
		git push origin main; \
	)

# ArgoCD app will be watching for changes in the local Gogs 42snippets_inception repo
42nginx:
	kubectl apply -f argoapp-42nginx.yaml
	#kubectl port-forward svc/nginx 8042:80
