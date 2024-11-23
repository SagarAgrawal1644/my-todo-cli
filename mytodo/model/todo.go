package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

type Tasks struct {
	Tasks []Todo `json:"tasks"`
}

type Todo struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due-date"`
	Completed   bool      `json:"completed"`
}

//will use a json file to store the todo list

// create an instance for a new task
func NewTodo(description string, dueDate time.Time) *Todo {
	return &Todo{
		Description: description,
		DueDate:     dueDate,
		Completed:   false,
	}
}

// array to store the tasks - add the functionality to load these from the file
var Todos []Todo

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

	//using simpletable package from alexeyco

	//create a new table
	table := simpletable.New()

	//defining the headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "Task Id"},
			{Text: "Description"},
			{Text: "Due Date"},
			{Text: "Completed"},
		},
	}

	// Add rows to the body
	table.Body = &simpletable.Body{
		Cells: [][]*simpletable.Cell{
			{
				{Text: "1"},
				{Text: "Learn Golang"},
				{Text: "added 5 days ago"},
				{Text: "pending"},
			},
		},
	}

	// table footer
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Text: "Total", Span: 2},
			{Text: "1 Task"},
		},
	}

	table.SetStyle(simpletable.StyleMarkdown)

	table.Print()

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
