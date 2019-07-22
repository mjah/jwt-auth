package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd consists of commands available to the application.
var RootCmd = &cobra.Command{
	Use:   "jwt-auth",
	Short: "A microservice to handle user authentication.",
	Long: `JWT Authentication Microservice

A microservice to handle user authentication.

More information available at the repository:
  https://github.com/mjah/jwt-auth`,
}

// Execute is the entry point to Cobra.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Read in config from file
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file `path` (default is $HOME/.jwt-auth.yaml)")

	// Set the default database url
	viper.SetDefault("database_url", "postgres://postgres:postgres@localhost:5432/jwt-auth")

	// Read in environment variables that match
	viper.AutomaticEnv()
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".jwt-auth" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".jwt-auth")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
