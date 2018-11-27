package cmd

import (
	"fmt"
	"strings"

	"github.com/mihirkelkar/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list",
	Long: `A Fast and Flexible Task Manger built with
                love by mihirkelkar and friends in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong while adding task")
		}
	},
}

// Add this to the root command
func init() {
	RootCmd.AddCommand(addCmd)
}
