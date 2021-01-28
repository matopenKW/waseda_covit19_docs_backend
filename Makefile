# Create Table
execute-sql:
ifeq ($(sqlname),)
	@echo "Please specify SQL file name"
	@echo "	$ make create-table sqlname=<sql file name>.sql"
else
	docker exec -it gowebprog_postgres /bin/bash -c "PGPASSWORD=gwp psql gwp -U gwp -f var/local/${sqlname}.sql"
endif
