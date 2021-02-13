#bin/bash
echo y | migrate -path db -database ${DATABASE_URL} $1 $2