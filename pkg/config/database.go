package config

import (
	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
	dbQuery    *dao.Query
}

// global variable for the connection instance
var connectionInstance *Database

func (database *Database) GetDBConnection() *gorm.DB {
	return database.connection
}

func (database *Database) GetDBQuery() *dao.Query {
	return database.dbQuery
}

func DbInstance() *Database {
	return connectionInstance
}

func NewDBConnection() (*Database, error) {
	dsn := "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// command: ~/go/bin/gentool -dsn "postgresql://postgres:postgres@localhost:5432/accounting_ledger?connect_timeout=10&sslmode=disable" -db postgres

	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return &Database{}, err
	}

	dbQuery := dao.Use(dbInstance)

	connectionInstance = &Database{
		connection: dbInstance,
		dbQuery:    dbQuery,
	}

	return connectionInstance, nil

}
