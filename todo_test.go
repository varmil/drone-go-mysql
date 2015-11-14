package main

import (
	"testing"
)

func TestTodo(t *testing.T) {
	db = connect("mysql", "root@tcp(127.0.0.1:3306)/todo")

	todos := []Todo{
		Todo{Title: "download drone"},
		Todo{Title: "setup continuous integration"},
		Todo{Title: "profit"},
	}

	// test saving to the database
	for _, todo := range todos {
		err := Save(&todo)
		if err != nil {
			t.Errorf("Error saving todo. %s", err)
		}
		if todo.ID == 0 {
			t.Errorf("Wanted todo id assigned, got 0")
		}
	}

	// test getting a list from the database
	list, err := List()
	if err != nil {
		t.Errorf("Error listing todo items. %s", err)
	}
	if len(list) != 3 {
		t.Errorf("Wanted %d items in list, got %d", len(todos), len(list))
	}
	for i, todo := range list {
		if todo.Title != todos[i].Title {
			t.Errorf("Wanted todo %s, got %s", todos[i].Title, todo.Title)
		}
	}

	// test deleting items from the database
	for _, todo := range list {
		err := Delete(todo.ID)
		if err != nil {
			t.Errorf("Error deleting todo %d. %s", todo.ID, err)
		}
	}
	list, _ = List()
	if len(list) != 0 {
		t.Errorf("Wanted 0 items in list, got %d", len(list))
	}
}
