package model

import (
	"fmt"

	config "github.com/israel-duff/ledger-system/pkg/config"
)

var db *config.Database

type UserModel struct {
	BaseModel
	EmailAddress string
	PhoneNumber  string
}

func MigrateUserTable() {
	fmt.Print("Auto migrating...")
	db = config.DB
	db.GetDBConnection().AutoMigrate(&UserModel{})
}

func (user *UserModel) Create() (*UserModel, error) {
	if db == nil {
		return nil, fmt.Errorf("Nil DB OBJECt")
	}
	result := db.GetDBConnection().Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Printf("Rows Affected %d", result.RowsAffected)

	return user, nil

}
