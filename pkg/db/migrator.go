package db

import (
	"embed"

	"github.com/pressly/goose/v3"
	"gorm.io/gorm"

	"fmt"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func RunMigrationUp(dbInstance *gorm.DB) {
	// var db *sql.DB
	db, err := dbInstance.DB()

	if err != nil {
		panic(err)
	}

	// setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// run app
}
