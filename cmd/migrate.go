package cmd

import (
	"github.com/mjah/jwt-auth/database"
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration.",
	Long:  `Run database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := database.Connect(); err != nil {
			logger.Log().Fatal("Could not connect to database. ", err)
		}

		if err := database.Migrate(); err != nil {
			logger.Log().Fatal("Could not migrate database. ", err)
		}
	},
}
