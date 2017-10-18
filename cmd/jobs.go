package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
// jobsCmd for manipulating jobs in the workspace
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "list, create, show, update or delete jobs",
	Long:  `list, create, show, update or delete jobs`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("jobs: list, show, create, update or delete")

	},
}

func init() {
	RootCmd.AddCommand(jobsCmd)

}
