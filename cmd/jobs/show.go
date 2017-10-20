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

var JobsShowCmd = &cobra.Command{
	Use:   "show [job_id]",
	Short: "show the given job",
	Long:  "show the given job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := utils.GetAPIToken()
		client := api.New(apiToken)
		job := client.GetJobByID(args[0])

		var status string
		if job.IsActive == true {
			status = fmt.Sprint(aurora.Green("active"))
		} else {
			status = fmt.Sprint(aurora.Red("incative"))
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		table.SetAutoMergeCells(true)
		table.Append([]string{"Name", job.Name})
		table.Append([]string{"ID", job.ID})
		table.Append([]string{"Status", status})
		table.Append([]string{"Cron Rule", job.CronRule})
		table.Append([]string{"Start Time", job.StartTime})
		table.Append([]string{"Timezone", job.Timezone})
		table.Append([]string{"Tag", job.Tag})
		table.Append([]string{"Action Type", job.ActionType})
		table.Append([]string{"Action Method", job.ActionMethod})
		table.Append([]string{"Action URL", job.ActionURL})
		for k, v := range job.ActionHeaders {
			table.Append([]string{"Action Headers", k + ": " + v})
		}
		// "action_query_parameters": {},
		// "action_headers": {},
		// "action_body": {},

		table.Render()
	},
}
