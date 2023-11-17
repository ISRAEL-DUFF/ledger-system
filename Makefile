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

test: 
	@go test ./... -v