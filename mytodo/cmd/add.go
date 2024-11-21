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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task.",
	Long: `Add a new task to your todo list.

	usage: mytodo add [description of task]`,
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	fmt.Println("add command called (ready to place your logic) ")

	if len(args) == 0 {
		fmt.Println("Please provide a task description.")
		return
	}

	newTodo := model.NewTodo(args[0], time.Now().Add(48*time.Hour))

	newTodo.Id = len(model.Todos) + 1

	model.AddTodo(*newTodo)

	fmt.Printf("Task '%s' added with due date %s\n", newTodo.Description, newTodo.DueDate.Format("2006-01-02"))

}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
