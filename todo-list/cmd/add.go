/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"mirzaadr/todo-cli/services"
	"mirzaadr/todo-cli/store"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "A command to add task to your todo list",
	Long:  `A command that will add a new task to your current list`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		todos := services.Todos{}

		storage := store.NewStore[services.Todos]("todos.json")
		storage.Load(&todos)

		title := strings.Join(args, " ")
		todos.Add(title)

		storage.Save(todos)
		// fmt.Printf("add called %s", title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
