package jobs

import (
	"fmt"

	"github.com/spf13/cobra"

	"../../api"
	"../../utils"
)

var JobsDisableCmd = &cobra.Command{
	Use:   "disable [job_id]",
	Short: "disable the given job",
	Long:  "disable the given job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := utils.GetAPIToken()
		client := api.New(apiToken)
		job := client.GetJobByID(args[0])

		job.IsActive = false
		_ = client.UpdateJob(job)

		fmt.Println("Success")

	},
}
