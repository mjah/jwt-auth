package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the authentication server.",
	Long:  `Run the authentication server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running authentication server...")
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrations...")
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(migrateCmd)
}
