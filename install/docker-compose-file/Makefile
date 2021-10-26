docker-compose-up-current-user:
	export CURRENT_UID=1000:1000 
	docker-compose up --force-recreate --abort-on-container-exit --remove-orphans
	# docker-compose up --force-recreate --renew-anon-volumes --build --abort-on-container-exit --remove-orphans
test-petstore:
	curl -v -H Host:petstore http://localhost:8080/petstore
