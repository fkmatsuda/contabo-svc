package cmd

import (
	"github.com/fkmatsuda/contabo-svc/lib/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Controls the Contabo VPS snapshot",
	Long: `The snapshot utility can create, delete, and list snapshots of the Contabo VPS.
	See create, delete, and list subcommands for more information.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Add autconfigh subcommand
	rootCmd.AddCommand(config.ConfigCmd)
	// Add snapshot management subcommands
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.contabo")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
