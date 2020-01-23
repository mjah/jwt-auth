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
	viper.SetDefault("postgres.database", "jwt-auth")
	viper.SetDefault("token.refresh_token_expires", "1d")
	viper.SetDefault("token.refresh_token_expires_extended", "1y")

	// Read binded and matching environment variables
	viper.AutomaticEnv()
	viper.BindEnv("token.private_key_path", "JA_PRIVATE_KEY_PATH")
	viper.BindEnv("postgres.host", "JA_PSQL_HOST")
	viper.BindEnv("postgres.port", "JA_PSQL_PORT")
	viper.BindEnv("postgres.username", "JA_PSQL_USERNAME")
	viper.BindEnv("postgres.password", "JA_PSQL_PASSWORD")
	viper.BindEnv("postgres.database", "JA_PSQL_DATABASE")
	viper.BindEnv("email.smtp_host", "JA_EMAIL_SMTP_HOST")
	viper.BindEnv("email.smtp_port", "JA_EMAIL_SMTP_PORT")
	viper.BindEnv("email.smtp_username", "JA_EMAIL_SMTP_USERNAME")
	viper.BindEnv("email.smtp_password", "JA_EMAIL_SMTP_PASSWORD")
	viper.BindEnv("email.from_address", "JA_EMAIL_FROM_ADDRESS")
	viper.BindEnv("email.test_receipient", "JA_EMAIL_TEST_RECEIPIENT")
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
