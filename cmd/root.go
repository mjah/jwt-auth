package cmd

import (
	"fmt"
	"os"

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

// Execute is the entry point to Cobra.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
