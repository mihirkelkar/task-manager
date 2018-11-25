package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Long: `A Fast and Flexible Task Manger built with
                love by mihirkelkar and friends in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		task := strings.Join(args, " ")
		fmt.Printf("Added %s to your task list.\n", task)
	},
}

// Add this to the root command
func init() {
	RootCmd.AddCommand(addCmd)
}
