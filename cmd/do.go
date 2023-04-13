package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed with a number",
	Run: func(cmd *cobra.Command, args []string) {
		var taskIds []int

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Argument \"%s\" is not a number\n", arg)
				continue
			}
			taskIds = append(taskIds, id)
		}

		fmt.Println(taskIds)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
