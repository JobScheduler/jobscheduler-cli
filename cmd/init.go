package cmd

import (
	"fmt"

	"../utils"
	"github.com/spf13/cobra"
)

// initCmd for associating with a workspace
var initCmd = &cobra.Command{
	Use:   "init [api_token]",
	Short: "Initialize the CLI",
	Long:  `Initialize the CLI`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := utils.SetAPIToken(args[0])
		// TODO: check apiToken validity
		utils.FatalIfErr(err)
		fmt.Println("Success!")

	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
