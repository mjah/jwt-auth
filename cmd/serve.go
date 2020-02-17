package cmd

import (
	"github.com/mjah/jwt-auth/api"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run authentication server.",
	Long:  `Run authentication server.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := database.Connect(); err != nil {
			logger.Log().Fatal(err)
		}
		database.Migrate()
		jwt.Setup()
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve()
	},
}
