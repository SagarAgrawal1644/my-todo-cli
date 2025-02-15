/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"mytodo/model"

	"github.com/spf13/cobra"
)

var dueDateInput string //variable to store the user due date

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task.",
	Long: `Add a new task to your todo list.

	usage: mytodo add [description of task]`,
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	//fmt.Println("add command called (ready to place your logic) ")

	if len(args) == 0 {
		fmt.Println("Please provide a task description.")
		return
	}

	taskDescription := args[0]

	//Parsing the user provided due date
	var dueDate time.Time
	var err error
	if dueDateInput != "" {
		dueDate, err = time.Parse("2006-01-02", dueDateInput) //user ro provide the date in yyyy-mm-dd format
		if err != nil {
			fmt.Println("Invalid date format. Please use yyyy-mm-dd.")
			return
		}
	} else {
		//defalut date if no date added using the -d flag
		dueDate = time.Now().Add(48 * time.Hour)
	}

	// call the AddTodo function to add the task
	if err := model.AddTodo(args[0], args[1]); err != nil {
		fmt.Println("Error adding Task: %v \n", err)
		return
	}
}

func init() {

	// flag to accept due date from the user
	addCmd.Flags().StringVarP(&dueDateInput, "due-date", "d", "", "Set custom due date in YYYY-MM-DD format")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
