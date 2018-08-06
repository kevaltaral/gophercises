package cmd

import (
	"fmt"

	"github.com/kevaltaral/gophercises/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command which lists all of your tasks.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:")
			return
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! ")
			return
		}
		//fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
