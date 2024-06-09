package main

import (
	"fmt"
	"os"
)

func main() {

	todos := []string{"first task", "second task", "third task"}

	// If no params are passed, list the todos
	if len(os.Args) == 1 {
		if len(todos) == 0 {
			fmt.Println("You have no todos")
			return
		} else {
			listTodos(todos)

		}
	}

}

func listTodos(todos []string) {
	for i, todo := range todos {
		fmt.Println(i+1, todo)
	}
}
