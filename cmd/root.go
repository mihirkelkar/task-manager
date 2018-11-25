package cmd

import "github.com/spf13/cobra"

//RootCmd := This is how you crate a CLI command using cobra`
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task manager is super fast custom built CLI task manager",
	Long: `A Fast and Flexible Task Manger built with
                love by mihirkelkar and friends in Go`,
}
