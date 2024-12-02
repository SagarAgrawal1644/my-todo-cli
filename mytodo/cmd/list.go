/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"mytodo/model"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display todo list.",
	Long: `The list command display your todo list with all the tasks and their status.

	usage: mytodo list`,
	Run: runList,
}

func runList(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
	//items to display
	displayItems, _ := model.ListTodos()

	//logic to display the items in a tabular form
	table := simpletable.New()

	//add headers
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "ID"},
			{Text: "Description"},
			{Text: "Due Date"},
			{Text: "Completed"},
		},
	}

	//adding rows for eac todo item
	for _, todo := range displayItems {
		// Format due date and completed status
		dueDate := todo.DueDate.Format("2006-01-02")
		completed := "No"
		if todo.Completed {
			completed = "Yes"
		}

		row := []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", todo.ID)},
			{Text: todo.Description},
			{Text: dueDate},
			{Text: completed},
		}

		table.Body.Cells == append(table.Body.Cells, row)

	}

	//add a footer if there are no todo items
	if len(displayItems) == 0 {
		table.Footer = &simpletable.Footer{
			Cells: []*simpletable.Cell{
				{Text: "No tasks found", Span: 4, Align: simpletable.AlignCenter},
			},
		}
	}

	//print table
	table.setStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
