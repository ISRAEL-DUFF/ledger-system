package main

import (
	"fmt"

	config "github.com/israel-duff/ledger-system/pkg/config"

	// models "github.com/israel-duff/ledger-system/pkg/models"
	"github.com/israel-duff/ledger-system/pkg/db"
)

func main() {
	// config.InitDatabaseConnection()
	// models.MigrateUserTable()
	fmt.Println("Hello world")

	// user := models.UserModel{
	// 	EmailAddress: "email1@gmail.com",
	// 	PhoneNumber:  "09028473643",
	// }

	// user.Create()

	// fmt.Printf("Use ID: %s", user.ID)

	databaseObject, err := config.NewDBConnection()

	if err != nil {
		panic("Could not connect to database!!!")
	}

	db.RunMigrationUp(databaseObject.GetDBConnection())

}
