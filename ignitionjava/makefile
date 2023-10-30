run:
	kubectl apply -f ignition-deployment.yaml
	kubectl apply -f ignition-service.yaml
	kubectl apply -f openldap-deployment.yaml
	kubectl apply -f openldap-service.yaml
	ldapadd -H ldap://localhost:30389 -x -D "cn=admin,dc=example,dc=org" -w admin -f base.ldif
	ldapadd -H ldap://localhost:30389 -x -D "cn=admin,dc=example,dc=org" -w admin -f admin.ldif
	ldapsearch -H ldap://localhost:30389 -x -D "cn=admin,dc=example,dc=org" -w admin -b "dc=example,dc=org" "(objectClass=*)"

clean:
	kubectl delete -f ignition-deployment.yaml
	kubectl delete -f ignition-service.yaml
	kubectl delete -f openldap-deployment.yaml
	kubectl delete -f openldap-service.yaml

#kubectl port-forward svc/ignition-service 8088:8088
