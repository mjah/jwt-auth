package cmd

import (
	"github.com/mjah/jwt-auth/api"
	"github.com/mjah/jwt-auth/auth"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/queue"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run authentication server.",
	Long:  `Run authentication server.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		database.Migrate()
		queue.Connect()
		auth.LoadPrivateKey()
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve()
	},
}
