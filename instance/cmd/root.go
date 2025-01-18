package cmd

import (
	"github.com/fkmatsuda/contabo-svc/lib/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "instance",
	Short: "Manage Contabo VPS instances",
	Long: `The instance utility can list VPS instances
	See list subcommand for more information.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add autconfigh subcommand
	rootCmd.AddCommand(config.ConfigCmd)
	// Add instance management subcommands
	rootCmd.AddCommand(listCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.contabo")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
