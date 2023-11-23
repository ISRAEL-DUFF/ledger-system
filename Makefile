# include .$(PWD)/.env
include .env
export

dbHost = $(DB_HOST)
dbPort = $(DB_PORT)
dbUser = $(DB_USER)
dbPassword = $(DB_PASSWORD)
dbName = $(DB_NAME)
timeZone = $(TZ)

dbDsn = "postgresql://${dbUser}:${dbPassword}@${dbHost}:${dbPort}/${dbName}?connect_timeout=10&sslmode=disable"


build:
	@cd cmd/ledger-system; go build -o ../../bin/ledger

run: build
	@./bin/ledger

builddriver:
	@cd cmd/test-drivers; go build -o ../../bin/testdriver

rundriver: builddriver
	@./bin/testdriver

buildresetdb:
	@cd cmd/db-reset; go build -o ../../bin/dbreset

dbreset: buildresetdb
	@./bin/dbreset

db-migrate:
	@~/go/bin/goose -dir pkg/db/migrations postgres "host=${dbHost} user=${dbUser} password=${dbPassword} dbname=${dbName} port=${dbPort} sslmode=disable TimeZone=${timeZone}" up

generate-models: db-migrate
	@~/go/bin/gentool -dsn ${dbDsn} -db postgres -outPath "pkg/db/dao"

test: 
	@go test ./... -v