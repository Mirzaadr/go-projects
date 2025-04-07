/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mirzaadr/todo-cli/services"
	"mirzaadr/todo-cli/store"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a task from the list",
	Long:  `this will delete the specified task from your todo list`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("index must be an integer")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		todos := services.Todos{}

		storage := store.NewStore[services.Todos]("todos.json")
		storage.Load(&todos)

		idx, _ := strconv.Atoi(args[0])
		todos.Delete(idx)

		storage.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
