package config

import (
	"fmt"
	"log"
	"os"

	"github.com/israel-duff/ledger-system/pkg/db/dao"
	"github.com/joho/godotenv"
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

type DBConnConfig struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DSN        string
}

func getDBEnv() (*DBConnConfig, error) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	timeZone := os.Getenv("TZ")

	config := &DBConnConfig{
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DSN:        fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", dbHost, dbUser, dbPassword, dbName, dbPort, timeZone),
	}

	return config, nil

}

func NewDBConnection() (*Database, error) {
	// dsn := "host=localhost user=postgres password=password dbname=accounting_ledger port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dbConfig, err := getDBEnv()

	if err != nil {
		return &Database{}, err
	}

	dsn := dbConfig.DSN
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
