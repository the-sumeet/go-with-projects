package main

import (
	"fmt"
	"os"
	"time"
)

type task struct {
	name string
	time time.Time
}

func main() {

	todos := []task{
		task{name: "Buy milk", time: time.Now()},
		task{name: "Buy bread", time: time.Now()},
	}

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

func listTodos(todos []task) {
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo.name)
	}
}
