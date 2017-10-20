package jobs

import (
	"fmt"

	"github.com/spf13/cobra"

	"../../api"
	"../../utils"
)

var JobsEnableCmd = &cobra.Command{
	Use:   "enable [job_id]",
	Short: "enable the given job",
	Long:  "enable the given job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := utils.GetAPIToken()
		client := api.New(apiToken)
		job := client.GetJobByID(args[0])

		job.IsActive = true
		_ = client.UpdateJob(job)

		fmt.Println("Success")

	},
}
