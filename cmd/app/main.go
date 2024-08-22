package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	database "simple-terminal-TODOList/internal/database"
	"strconv"
)

type Todo struct {
	ID   int
	Task string
	Done bool
}

var db *sql.DB

func main() {
	db = database.ConnectDB()
	fmt.Println("Подключение к БД прошло успешно")

	defer db.Close()

	err := database.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	// Обработка команд из аргументов командной строки
	if len(os.Args) < 2 {
		fmt.Println("Использование: todo <команда> [аргументы]")
		fmt.Println("Команды: add, list, done, delete")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Использование: todo add <задача>")
			return
		}
		addTodo(os.Args[2:])
	case "list":
		listTodos()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Использование: todo done <ID задачи>")
			return
		}
		markTodoDone(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Использование: todo delete <ID задачи>")
			return
		}
		deleteTodo(os.Args[2])
	default:
		fmt.Println("Неизвестная команда:", os.Args[1])
	}
}

func addTodo(task []string) {
	query := `INSERT INTO todos(task, done) VALUES(?, ?)`
	_, err := db.Exec(query, fmt.Sprintf("%s", task), false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Задача добавлена:", fmt.Sprintf("%s", task))
}

func listTodos() {
	rows, err := db.Query("SELECT id, task, done FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Список задач:")
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Task, &todo.Done)
		if err != nil {
			log.Fatal(err)
		}
		status := " "
		if todo.Done {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Task)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func markTodoDone(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Неверный ID:", idStr)
		return
	}

	query := `UPDATE todos SET done = true WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Задача с таким ID не найдена:", id)
	} else {
		fmt.Println("Задача выполнена:", id)
	}
}

func deleteTodo(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Неверный ID:", idStr)
		return
	}

	query := `DELETE FROM todos WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Println("Задача с таким ID не найдена:", id)
	} else {
		fmt.Println("Задача удалена:", id)
	}
}
