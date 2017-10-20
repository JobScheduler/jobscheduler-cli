package jobs

import (
	"github.com/spf13/cobra"
)

var JobsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update the given job",
	Long:  "update the given job",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
