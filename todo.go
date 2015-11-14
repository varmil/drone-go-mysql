package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	ID    int64  // Unique identifier
	Title string // Description
}

// Database connection
var db *sql.DB

// Save saves the given Todo in the database.
func Save(todo *Todo) error {
	res, err := db.Exec("INSERT INTO todos VALUES (null, ?)", todo.Title)
	if err != nil {
		return err
	}
	todo.ID, err = res.LastInsertId()
	return err
}

// All returns the list of all the Tasks in the database.
func List() ([]*Todo, error) {
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		todo := &Todo{}
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
		)
		if err != nil {
			break
		}
		todos = append(todos, todo)
	}
	return todos, err
}

// Delete deltes the Todo with the given id in the database.
func Delete(id int64) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

// connect is a helper function that creates the database
// connection and creates required tables.
func connect(driver, datasource string) *sql.DB {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}
	return db
}

const schema = `
CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTO_INCREMENT, 
	title VARCHAR(2000)
);
`
