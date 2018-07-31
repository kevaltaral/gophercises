package cmd

import (
	"fmt"
	"strings"

	"github.com/kevaltaral/gophercises/task/db"

	"github.com/spf13/cobra"
)

// addCmd represents the add command which adds tasks to your task list
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Print("Not able to add tasks ")
		}
		fmt.Printf("Added \"%s\"\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
