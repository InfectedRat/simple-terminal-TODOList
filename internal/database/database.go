package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
	DBpath := "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db"
	db, err := sql.Open("sqlite3", DBpath)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("База не доступна: %v", err)
	}
	return db
}

func CreateTable(db *sql.DB) error {

	query := `CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT,
        done BOOLEAN);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("Таблица notes успешно создана или уже существует.")
	return err
}
