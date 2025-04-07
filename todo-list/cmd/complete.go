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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <taskid>",
	Short: "mark task as complete",
	Long:  `A function to change a task from the list to complete`,
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
		todos.Complete(idx)

		storage.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
