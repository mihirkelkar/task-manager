package cmd

import (
	"fmt"
	"os"

	"github.com/mihirkelkar/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks to complete")
			return
		}
		fmt.Println("You have the following tasks to complete")
		fmt.Println("========================================")
		for index, task := range tasks {
			fmt.Printf("%d. %s\n", index+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
