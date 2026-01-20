package data

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const (
	dbDriver = "sqlite"
	dbName   = "data/notaDB"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado ao SQLite com sucesso!")
	return db, nil
}
