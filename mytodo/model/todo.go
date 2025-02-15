package model

import (
	"fmt"
	"mytodo/mytodo/database"
	"strconv"
	"time"
)

type Todo struct {
	Id          int
	Description string
	DueDate     time.Time
	Completed   bool
}

// add the new task to the array
func AddTodo(description string, dueDate time.Time) error {
	query := "INSERT INTO tasks (description, due_date, completed) VALUES (?, ?, ?,)"
	_, err := database.DB.Exec(query, description, dueDate.Format("2006-01-02"), false)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}
	fmt.Println("Task added successfully!")
	return nil

}

// list all the tasks
func ListTodos() ([]Todo, error) {
	query := "SELECT id, description, due_date, completed FROM tasks"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		var dueDate string
		err := rows.Scan(&todo.ID, &todo.Description, &dueDate, &todo.Completed)
		if err != nil {

		}
		todo.DueDate, _ = time.Parse("2006-01-02", dueDate)
		todos = append(todos, todo)
	}

	return todos, nil
}

// mark a task as done
func MarkAsCompleted(id int) {

	query := "SELECT id, description, due_date, completed FROM tasks WHERE id = ?"
	row, err := DB.QueryRow(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to search the task: %w", err)
	}

	var task Todo

	err := row.Scan(&task.ID, &task.Description, &task.dueDate, &task.Completed)
	
	if err == sql.ErrNoRows {
		fmt.Println("Enter a valid task ID.")
		return
	}
	else if err != nil {
		fmt.Errorf("failed to scan row: %w", err)
		return
	}

	// check if the task is already completed
	if task.Completed {
		fmt.Println("Task is already completed.")
		return 
	}

	// update the task by a new query
	updateQuery := "UPDATE tasks SET completed = ? WHERE id = ?"
	_, err = db.Exec(updateQuery, true, id)
	if err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	fmt.Printf("Task with ID %d marked as completed.\n", id)

}
