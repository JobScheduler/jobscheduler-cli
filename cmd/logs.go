package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logsCmd for inspecting logs in the workspace
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Inspect logs",
	Long:  `Inspect logs`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Inspect logs")

	},
}

func init() {
	RootCmd.AddCommand(logsCmd)
}
