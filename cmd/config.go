package cmd

import (
	"strings"

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
	viper.SetDefault("log_email", true)
	viper.SetDefault("account.password_cost", 11)
	viper.SetDefault("account.confirm_token_expires", "24h00m")
	viper.SetDefault("account.reset_password_token_expires", "1h00m")
	viper.SetDefault("roles.define", []string{"admin", "member", "guest"})
	viper.SetDefault("roles.default", "guest")
	viper.SetDefault("serve.host", "localhost")
	viper.SetDefault("serve.port", 9096)
	viper.SetDefault("cors.allow_all_origins", true)
	viper.SetDefault("cors.allow_origins", []string{"http://localhost:8080"})
	viper.SetDefault("cors.allow_credentials", false)
	viper.SetDefault("token.issuer", "jwt-auth")
	viper.SetDefault("token.access_token_expires", "5m")
	viper.SetDefault("token.refresh_token_expires", "8h00m")
	viper.SetDefault("token.refresh_token_expires_extended", "8760h00m")
	viper.SetDefault("postgres.host", "localhost")
	viper.SetDefault("postgres.port", 5432)
	viper.SetDefault("postgres.sslmode", "disable")
	viper.SetDefault("postgres.username", "postgres")
	viper.SetDefault("postgres.password", "postgres")
	viper.SetDefault("postgres.database", "jwt-auth")
	viper.SetDefault("amqp.host", "localhost")
	viper.SetDefault("amqp.port", 5672)
	viper.SetDefault("amqp.username", "guest")
	viper.SetDefault("amqp.password", "guest")

	// Read environment variables
	viper.SetEnvPrefix("JA")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			logger.Log().Fatal("Could not read config. ", err)
		}
	}
}
