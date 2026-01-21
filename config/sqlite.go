package config

import (
	"ProjetoNota/schemas"
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	dbPath = "./db/main.db"
)

func InitializeSQLite() (*gorm.DB, error) {

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		fmt.Println("Database file doesn't exist")
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		_ = file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Nota{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
