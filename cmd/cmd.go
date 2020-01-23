package cmd

import (
	"fmt"
	"os"

	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/email"
	"github.com/spf13/cobra"
)

// RootCmd consists of commands available to the application.
var RootCmd = &cobra.Command{
	Use:   "jwt-auth",
	Short: "A microservice to handle user authentication.",
	Long: `JWT Authentication Microservice

A microservice to handle user authentication.

More information available at the repository:
  https://github.com/mjah/jwt-auth`,
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run authentication server.",
	Long:  `Run authentication server.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		database.Migrate()
	},
	Run: func(cmd *cobra.Command, args []string) {
		auth.Run()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration.",
	Long:  `Run database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Migrate()
	},
}

var testEmailCmd = &cobra.Command{
	Use:   "test-email",
	Short: "Send a test email.",
	Long:  `Send a test email.`,
	Run: func(cmd *cobra.Command, args []string) {
		email.Test()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(migrateCmd)
	RootCmd.AddCommand(testEmailCmd)
}

// Execute is the entry point to Cobra.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
