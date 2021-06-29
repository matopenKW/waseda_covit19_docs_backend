# init
init:
	docker-compose up -d app
	docker exec -it local_app /bin/sh -c "go get"
	docker-compose stop app

db-run:
	docker-compose up -d db

server-run:
	docker-compose up -d app

migrate-up:
	docker-compose up -d migrate
	docker exec -it migrate /bin/bash -c "bash migrate.bash up"

migrate-down:
	docker-compose up -d migrate
	docker exec -it migrate /bin/bash -c "bash migrate.bash down"

migrate-force:
ifeq ($(ver),)
	@echo "Please specify version"
	@echo "	$ make migrate-force ver=<version>"
else
	docker-compose up -d migrate
	docker exec -it migrate /bin/bash -c "bash migrate.bash force ${ver}"
endif

create-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make create-sql sqlname=<sql file name>.sql"
else
	docker-compose up -d migrate
	docker exec -it migrate /bin/bash -c "migrate create -ext sql -dir db -seq ${sqlname}"
endif

run:
	make db-run
	make server-run
	docker exec -it local_app /bin/sh -c "go run main.go"

stop:
	docker-compose stop

execute-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make execute-sql sqlname=<sql file name>.sql"
else
	docker exec -it local_db /bin/bash -c "PGPASSWORD=gwp psql gwp -U gwp -f var/local/${sqlname}.sql"
endif

go-test:
	docker-compose up -d app
	docker exec -it local_app /bin/sh -c "go test ./..."

cmd-run:
	# docker build ./local/cmd/ -t local_cmd
	# docker run -v $(pwd):/go/src --env-file .env --rm -it --name local_cmd local_cmd
	docker-compose -f ./docker-compose.cmd.yaml up -d
	docker exec -it local_cmd /bin/bash

db-run-k8s:
	kubectl config use-context docker-desktop
	kubectl delete -k ./k8s/overlays/local/db/ --ignore-not-found=true
	kubectl apply -f ./k8s/overlays/local/db/pvc.yaml
	kubectl apply -k ./k8s/overlays/local/db/
