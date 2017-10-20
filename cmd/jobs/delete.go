package jobs

import (
	"github.com/spf13/cobra"
)

var JobsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the given job",
	Long:  "delete the given job",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
