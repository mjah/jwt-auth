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

	// Read in config from file
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file `path`")

	// Set the default database url
	viper.SetDefault("token.access_token_expires", "5m")
	viper.SetDefault("token.refresh_token_expires", "1d")
	viper.SetDefault("token.refresh_token_expires_extended", "1y")

	// Read in environment variables that match
	// viper.AutomaticEnv()
	viper.BindEnv("token.private_key_path", "JA_PRIVATE_KEY")
	viper.BindEnv("postgres.host", "JA_HOST")
	viper.BindEnv("postgres.port", "JA_PORT")
	viper.BindEnv("postgres.username", "JA_USERNAME")
	viper.BindEnv("postgres.password", "JA_PASSWORD")
	viper.BindEnv("postgres.database", "JA_DATABASE")
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
