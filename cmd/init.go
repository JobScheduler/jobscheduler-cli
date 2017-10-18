package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd for associating with a workspace
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI",
	Long:  `Initialize the CLI`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Initialize the CLI")

	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
