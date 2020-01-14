package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	// Read config from file
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file `path`")

	// Set default values
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.password", "jwt-auth")
	viper.SetDefault("token.refresh_token_expires", "1d")
	viper.SetDefault("token.refresh_token_expires_extended", "1y")

	// Read binded and matching environment variables
	viper.AutomaticEnv()
	viper.BindEnv("token.private_key_path", "JA_PRIVATE_KEY")
	viper.BindEnv("postgres.host", "JA_PSQL_HOST")
	viper.BindEnv("postgres.port", "JA_PSQL_PORT")
	viper.BindEnv("postgres.username", "JA_PSQL_USERNAME")
	viper.BindEnv("postgres.password", "JA_PSQL_PASSWORD")
	viper.BindEnv("postgres.database", "JA_PSQL_DATABASE")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config: ", err)
			os.Exit(1)
		}
	}
}
