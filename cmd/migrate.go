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
			logger.Log().Fatal("Failed to connect to database.")
		}
		database.Migrate()
	},
}
