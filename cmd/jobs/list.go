package jobs

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"../../api"
	"../../utils"
)

var JobsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list jobs",
	Long:  "list workspace's jobs",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := utils.GetAPIToken()
		client := api.New(apiToken)
		jobs := client.GetJobs()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Status", "ID", "Cron Rule", "Timezone"})

		for _, job := range jobs {
			var status string
			if job.IsActive == true {
				status = fmt.Sprint(aurora.Green("active"))
			} else {
				status = fmt.Sprint(aurora.Red("incative"))
			}
			table.Append([]string{job.Name, status, job.ID, job.CronRule, job.Timezone})
		}
		table.Render()

	},
}
