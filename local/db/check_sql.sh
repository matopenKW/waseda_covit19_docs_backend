#bin/bash

for f in ./sql/*; do
    SQL_FILE=$(basename ${f%.*})
    docker exec -it local_db /bin/bash -c "PGPASSWORD=gwp psql gwp -U gwp -f var/local/${SQL_FILE}.sql"

    docker exec -it local_db /bin/bash -c 'false; echo $? > /tmp/return-code'
    echo $(docker exec local_db /bin/cat /tmp/return-code)
done