/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"mytodo/model"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark task as completed",
	Long: `Mark your task as done or completed by its id

	usage: mytodo done [task id]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")

		if len(args) == 0 {
			fmt.Println("Enter the id of the task you want to delete {use the 'list' command to see to-do list")
			return
		}

		if len(model.Todos) == 0 {
			fmt.Println("Add some task to mark it as completed {use the 'add' command to add tasks}")
			return
		}

		id := args[0]
		model.MarkAsCompleted(id)

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
