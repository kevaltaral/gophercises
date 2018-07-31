package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI Task manager",
}
