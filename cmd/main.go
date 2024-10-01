package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NoistNT/go-cli-todo/internal/todos"
	helpers "github.com/NoistNT/go-cli-todo/pkg"
)

func main() {
	const (
		titleMsg      = "What should we write in the todo?: "
		errorIDMsg    = "Error reading ID"
		errorTitleMsg = "Error reading title"
	)

	todos.Init()
	helpers.PrintCommands()
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := helpers.GetInput(reader, "[cmd]: ")
		helpers.ErrorHandler(err)

		switch cmd {
		case "add":
			title, err := helpers.GetInput(reader, titleMsg)
			helpers.ErrorHandler(err)
			todos.Add(title)

		case "list":
			todos.List()

		case "update":
			id, err := helpers.GetInputInt(reader, "Enter todo ID to update: ")
			helpers.ErrorHandler(err)
			title, err := helpers.GetInput(reader, titleMsg)
			helpers.ErrorHandler(err)
			todos.Update(id, title)

		case "done":
			id, err := helpers.GetInputInt(reader, "Enter todo ID to mark as done: ")
			helpers.ErrorHandler(err)
			todos.Done(id)

		case "remove":
			id, err := helpers.GetInputInt(reader, "Enter todo ID to remove: ")
			helpers.ErrorHandler(err)
			todos.Remove(id)

		case "exit", "quit":
			fmt.Println("Program ended successfully")
			os.Exit(0)

		default:
			fmt.Println("Invalid command!")
			helpers.PrintCommands()
		}
	}
}
