package handler

import (
	"ProjetoNota/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Initialize() {
	db = config.GetSQLite()
}
