package cmd

import (
	"../utils"
	"github.com/spf13/cobra"
)

// initCmd for associating with a workspace
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI",
	Long:  `Initialize the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			return
		}

		err := utils.SetAPIToken(args[0])
		utils.FatalIfErr(err)

	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
