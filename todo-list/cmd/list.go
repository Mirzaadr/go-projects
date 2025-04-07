/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mirzaadr/todo-cli/services"
	"mirzaadr/todo-cli/store"

	"github.com/spf13/cobra"
)

var isAll *bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the todo list",
	Long:  `print all available todo as a table`,
	Run: func(cmd *cobra.Command, args []string) {
		todos := services.Todos{}

		storage := store.NewStore[services.Todos]("todos.json")
		storage.Load(&todos)

		todos.Print(*isAll)
		storage.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	isAll = listCmd.Flags().BoolP("all", "a", false, "print all info")
}
