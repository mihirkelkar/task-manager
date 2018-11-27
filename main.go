package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mihirkelkar/task/cmd"
	"github.com/mihirkelkar/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	db.AllTasks()
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
