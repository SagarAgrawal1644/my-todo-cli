package model

import (
	"fmt"
	"strconv"
	"time"
)

type Todo struct {
	Id          int
	Description string
	DueDate     time.Time
	Completed   bool
}

// array to store tasks
var Todos []Todo

// create an instance for a new task
func NewTodo(description string, dueDate time.Time) *Todo {
	return &Todo{
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
}

// add the new task to the array
func AddTodo(newTodo Todo) {
	Todos = append(Todos, newTodo)
}

// list all the tasks
func ListTodos() {
	if len(Todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("To-do List")
	//logic for iterating and printing all the tasks
	//wanna try to make it tabular using more packages

}

// mark a task as done
func MarkAsCompleted(idstr string) {

	id, _ := strconv.Atoi(idstr)

	//search the item by its id and mark completed as true
	for _, item := range Todos {
		if item.Id == id {
			if item.Completed == true {
				fmt.Println("Task already completed.")
				return
			} else {
				item.Completed = true
				return
			}
		}
	}
	fmt.Println("Enter a valid ID {use the 'list' command to display your to-do list")
	return
}
