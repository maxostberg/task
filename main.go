package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maxostberg/task/cmd"
	"github.com/maxostberg/task/db"
)

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Could not find home directory: %s\n", err.Error())
		os.Exit(1)
	}

	dbPath := filepath.Join(dir, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}
