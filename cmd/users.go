package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// usersCmd for modifying in the workspace
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Add or remove users",
	Long:  "Add or remove users from workspace",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("modify users")

	},
}

func init() {
	RootCmd.AddCommand(usersCmd)
}
