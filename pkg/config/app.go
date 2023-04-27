package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func connectDB() {
	dsn := "host=localhost user=israel password=password dbname=account_ledger port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = dbInstance

}

func getDB() *gorm.DB {
	return db
}
