package cmd

import (
	"github.com/mjah/jwt-auth/api"
	"github.com/mjah/jwt-auth/auth/jwt"
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/email"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run authentication server.",
	Long:  `Run authentication server.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := database.Connect(); err != nil {
			logger.Log().Fatal("Could not connect to database. ", err)
		}

		if err := database.Migrate(); err != nil {
			logger.Log().Fatal("Could not migrate database. ", err)
		}

		if err := email.Setup(); err != nil {
			logger.Log().Fatal("Could not setup email. ", err)
		}

		if err := jwt.Setup(); err != nil {
			logger.Log().Fatal("Could not setup JWT. ", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve()
	},
}
