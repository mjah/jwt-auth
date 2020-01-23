package cmd

import (
	"github.com/mjah/jwt-auth/database"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration.",
	Long:  `Run database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Migrate()
	},
}
