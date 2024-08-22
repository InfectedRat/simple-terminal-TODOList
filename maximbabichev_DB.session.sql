SELECT * FROM accounts_tinkoff

func ConnectDB() *sql.DB {
	databasePath := "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db"

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	return db

}

func CreateTable(db *sql.DB) {

	query := `
		CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("Таблица notes успешно создана или уже существует.")
}