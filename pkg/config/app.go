package config

import (
	"fmt"
)

var (
	DB *Database
)

func InitDatabaseConnection() {
	fmt.Print("Creating DATABAse connection...")
	dbConn, err := NewDBConnection()

	if err != nil {
		panic(err)
	}

	DB = dbConn

}

// func ConnectDB() {
// 	dsn := "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	db = dbInstance

// }

// func GetDB() *gorm.DB {
// 	return db
// }
