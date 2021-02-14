# init
init:
	docker-compose up -d app
	docker exec -it local_app /bin/sh -c "go get"
	docker-compose stop app

db-run:
	docker-compose up -d db

server-run:
	docker-compose up -d app
	docker exec -it local_app /bin/sh -c "go run main.go"

run:
	make db-run
	make server-run

stop:
	docker-compose stop

execute-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make create-table sqlname=<sql file name>.sql"
else
	docker exec -it local_db /bin/bash -c "PGPASSWORD=gwp psql gwp -U gwp -f var/local/${sqlname}.sql"
endif
