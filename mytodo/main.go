/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"mytodo/cmd"
	"mytodo/database"
)

func main() {
	// Initialize the database
	database.InitDB()
	defer database.CloseDB()

	cmd.Execute()
}
