# init
init:
	docker-compose up -d
	docker exec -it local_app /bin/sh -c "go get"
	docker-compose stop app


db-run:
	docker-compose up -d db

server-run:
	docker-compose up -d app
	docker exec -it local_app /bin/sh -c "go run main.go"

migrate-up:
	docker-compose up -d migrate
	docker exec -it migrate /bin/bash -c "migrate -path db -database ${DATABASE_URL} up"
	docker-compose stop migrate

migrate-down:
	docker-compose up -d migrate
	echo y | migrate -path db/migrations -database ${DATABASE_URL} down
	docker-compose stop migrate

migrate-force:
ifeq ($(version),)
	@echo "Please specify version"
	@echo "	$ make migrate-force version=<version>"
else
	docker-compose up -d migrate
	migrate -path db -database ${DATABASE_URL} force ${version}
	docker-compose stop migrate
endif

create-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make create-sql sqlname=<sql file name>.sql"
else
	docker-compose up -d migrate
	migrate create -ext sql -dir db -seq ${sqlname}
	docker-compose stop migrate
endif

run:
	make db-run
	make server-run

stop:
	docker-compose stop

execute-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make execute-sql sqlname=<sql file name>.sql"
else
	docker exec -it local_db /bin/bash -c "PGPASSWORD=gwp psql gwp -U gwp -f var/local/${sqlname}.sql"
endif
