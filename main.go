package main

import (
	"fmt"
	"os"

	"github.com/mihirkelkar/task/cmd"
	"github.com/mihirkelkar/task/db"
)

func main() {
	err := db.Init("tasks.db")
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
