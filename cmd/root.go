package cmd

import (
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(migrateCmd)
	RootCmd.AddCommand(testEmailCmd)
}

// Execute is the entry point to Cobra.
func Execute() {
	logger.SetupLog(viper.GetString("log_level"))

	if err := RootCmd.Execute(); err != nil {
		// logger.Log().Fatal(err)
	}
}
