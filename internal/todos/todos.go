package todos

import (
	"fmt"

	helpers "github.com/NoistNT/go-cli-todo/pkg"
)

// Todo struct
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"status"`
}

// Todos struct
type Todos []Todo

// todosList is the list of todos
var todosList Todos

// Seed seeds the list with some todos
func Seed() {
	fmt.Println("Seeding todos...")
	todosList = Todos{
		{ID: 1, Title: "Buy milk", Done: false},
		{ID: 2, Title: "Buy eggs", Done: false},
		{ID: 3, Title: "Buy bread", Done: true},
		{ID: 4, Title: "Buy cheese", Done: true},
		{ID: 5, Title: "Buy water", Done: false},
	}
	fmt.Println("Todos seeded successfully.")
}

// NewTodo returns a new Todo
func NewTodo(title string) *Todo {
	return &Todo{ID: len(todosList) + 1, Title: title, Done: false}
}

// Add adds a new todo to the list
func Add(title string) {
	todosList = append(todosList, *NewTodo(title))
	fmt.Println("Todo added successfully.")
}

// List prints the list of todos
func List() {
	if len(todosList) == 0 {
		fmt.Println("Todo List is empty, try adding some todos using the 'add' command.")
		return
	}
	for _, todo := range todosList {
		fmt.Printf("[%s] %d: %s\n", doneIcon(todo.Done), todo.ID, todo.Title)
	}
}

// Update updates todo title with matching ID
func Update(id int, title string) {
	todo, _, err := find(id)
	helpers.ErrorHandler(err)

	todo.Title = title
	fmt.Printf("Todo with ID %d updated successfully.\n", id)
}

// Done updates todo done to true or false
func Done(id int) {
	todo, _, err := find(id)
	helpers.ErrorHandler(err)

	todo.Done = !todo.Done
	fmt.Printf("Todo with ID %d is now marked as %s.\n", id, doneIcon(todo.Done))
}

// Remove todo with matching ID
func Remove(id int) {
	_, i, err := find(id)
	if i < 0 {
		helpers.ErrorHandler(err)
		return
	}
	helpers.ErrorHandler(err)

	todosList = append(todosList[:i], todosList[i+1:]...)
	fmt.Printf("Todo with ID %d removed successfully.\n", id)
}

// Find finds todo with matching ID or error
func find(id int) (*Todo, int, error) {
	for i, todo := range todosList {
		if todo.ID == id {
			return &todosList[i], i, nil
		}
	}
	return nil, -1, fmt.Errorf("Todo with ID %d not found", id)
}

// DoneIcon returns ✅ or ❌
func doneIcon(done bool) string {
	var icon string
	if done {
		icon = "✅"
	} else {
		icon = "❌"
	}
	return icon
}
