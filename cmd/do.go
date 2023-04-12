package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed with a number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Args)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
