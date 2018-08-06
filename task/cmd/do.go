package cmd

import (
	"fmt"

	"strconv"

	"github.com/kevaltaral/gophercises/task/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command and Marks a task as complete
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:")
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			msg := fmt.Sprintf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)

			if err == nil {
				msg = fmt.Sprintf("Marked \"%d\" as completed.\n", id)

			}
			fmt.Printf(msg)
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
