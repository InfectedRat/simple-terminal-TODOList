package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	DBpath := "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db"
	db, err := sql.Open("sqlite3", DBpath)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("База не доступна", err)
	}
	return db
}
