package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

func main() {
	var todos *TodoManager

	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "manage a todo list"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "driver",
			Value: "mysql",
			Usage: "database driver name",
		},
		cli.StringFlag{
			Name:  "datasource",
			Value: "root@tcp(127.0.0.1:3306)/todo",
			Usage: "database configuration string",
		},
	}
	app.Before = func(c *cli.Context) error {
		var err error
		todos, err = NewTodoManager(
			c.GlobalString("driver"),
			c.GlobalString("datasource"),
		)
		return err
	}
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add a todo to the list",
			Action: func(c *cli.Context) {
				// creates a new todo using the title specified
				// in the command line input
				todo := &Todo{}
				todo.Title = c.Args().First()

				// save the todo to the datbase
				err := todos.Save(todo)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("added todo number %d\n", todo.ID)
			},
		},
		{
			Name:  "rm",
			Usage: "remove a todo from the list",
			Action: func(c *cli.Context) {
				// parses the todo id from the command line arg
				id, err := strconv.ParseInt(c.Args().First(), 10, 64)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// deletes the todo by id
				err = todos.Delete(id)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("deleted todo %d\n", id)
			},
		},
		{
			Name:  "ls",
			Usage: "lists all todos",
			Action: func(c *cli.Context) {
				// fetch all todo items from the database
				list, err := todos.List()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// print each to the console
				for _, todo := range list {
					fmt.Println(todo.ID, todo.Title)
				}
			},
		},
	}
	app.Run(os.Args)
}
