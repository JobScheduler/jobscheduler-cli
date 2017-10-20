package jobs

import (
	"github.com/spf13/cobra"
)

var JobsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a job",
	Long:  "Create a job in the workspace",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
