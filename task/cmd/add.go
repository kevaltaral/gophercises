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
		msg := "Not able to add tasks"
		_, err := db.CreateTask(task)
		if err == nil {
			msg = fmt.Sprintf("Added \"%s\"\n", task)
		}
		fmt.Print(msg)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
