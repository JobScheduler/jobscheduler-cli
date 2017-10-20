package jobs

import (
	"github.com/spf13/cobra"
)

var JobsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show the given job",
	Long:  "show the given job",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
