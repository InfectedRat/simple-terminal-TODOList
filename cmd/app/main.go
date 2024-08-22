package main

import (
	"fmt"
	database "simple-terminal-TODOList/internal/database"
)

func main() {
	database.ConnectDB()
	fmt.Println("Подключение к БД прошло успешно")
}
