package cmd

import (
	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	// Read config from file
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file `path`")

	// Set default values
	viper.SetDefault("environment", "development")
	viper.SetDefault("log_level", "debug")
	viper.SetDefault("serve.host", "localhost")
	viper.SetDefault("serve.port", 9096)
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.database", "jwt-auth")
	viper.SetDefault("amqp.host", "localhost")
	viper.SetDefault("amqp.port", 5672)
	viper.SetDefault("amqp.username", "guest")
	viper.SetDefault("amqp.password", "guest")
	viper.SetDefault("token.refresh_token_expires", "1d")
	viper.SetDefault("token.refresh_token_expires_extended", "1y")

	// Read binded and matching environment variables
	viper.AutomaticEnv()
	viper.BindEnv("environment", "JA_ENVIRONMENT")
	viper.BindEnv("log_level", "JA_LOG_LEVEL")
	viper.BindEnv("serve.host", "JA_SERVE_HOST")
	viper.BindEnv("serve.port", "JA_SERVE_PORT")
	viper.BindEnv("token.private_key_path", "JA_PRIVATE_KEY_PATH")
	viper.BindEnv("postgres.host", "JA_PSQL_HOST")
	viper.BindEnv("postgres.port", "JA_PSQL_PORT")
	viper.BindEnv("postgres.username", "JA_PSQL_USERNAME")
	viper.BindEnv("postgres.password", "JA_PSQL_PASSWORD")
	viper.BindEnv("postgres.database", "JA_PSQL_DATABASE")
	viper.BindEnv("amqp.host", "JA_AMQP_HOST")
	viper.BindEnv("amqp.port", "JA_AMQP_PORT")
	viper.BindEnv("amqp.username", "JA_AMQP_USERNAME")
	viper.BindEnv("amqp.password", "JA_AMQP_PASSWORD")
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
			logger.Log().Fatal("Can't read config: ", err)
		}
	}
}
