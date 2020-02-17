package cmd

import (
	"github.com/mjah/jwt-auth/email"
	"github.com/spf13/cobra"
)

var testEmailCmd = &cobra.Command{
	Use:   "test-email",
	Short: "Send a test email.",
	Long:  `Send a test email.`,
	Run: func(cmd *cobra.Command, args []string) {
		email.SendTestEmail()
	},
}
