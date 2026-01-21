package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeConfig() error {
	var err error
	db, err = InitializeSQLite()

	if err != nil {
		return err
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
