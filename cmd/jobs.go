package cmd

import (
	"./jobs"
	"github.com/spf13/cobra"
)

// jobsCmd for manipulating jobs in the workspace
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "list, create, show, update or delete jobs",
	Long:  `list, create, show, update or delete jobs`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(jobsCmd)
	jobsCmd.AddCommand(jobs.JobsCreateCmd)
	jobsCmd.AddCommand(jobs.JobsDeleteCmd)
	jobsCmd.AddCommand(jobs.JobsListCmd)
	jobsCmd.AddCommand(jobs.JobsShowCmd)
	jobsCmd.AddCommand(jobs.JobsUpdateCmd)
	jobsCmd.AddCommand(jobs.JobsEnableCmd)
	jobsCmd.AddCommand(jobs.JobsDisableCmd)
}
