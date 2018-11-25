package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
		var ids []int
		for _, nums := range args {
			id, err := strconv.Atoi(nums)
			if err != nil {
				fmt.Println("Failed to parse the arguement: ", nums)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
