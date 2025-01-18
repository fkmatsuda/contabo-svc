package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Contabo API authentication",
	Run: func(cmd *cobra.Command, args []string) {
		clientId, _ := cmd.Flags().GetString("client-id")
		clientSecret, _ := cmd.Flags().GetString("client-secret")
		apiUser, _ := cmd.Flags().GetString("api-user")
		apiPassword, _ := cmd.Flags().GetString("api-password")

		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configDir := filepath.Join(homeDir, ".contabo")

		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			err := os.Mkdir(configDir, 0755)
			if err != nil {
				panic(err)
			}
		}
		configFile := filepath.Join(configDir, "config.json")
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			_, err := os.Create(configFile)
			if err != nil {
				panic(err)
			}
		}

		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("$HOME/.contabo")

		viper.Set("auth.client_id", clientId)
		viper.Set("auth.client_secret", clientSecret)
		viper.Set("auth.api_user", apiUser)
		viper.Set("auth.api_password", apiPassword)

		err = viper.WriteConfig()
		if err != nil {
			panic(err)
		}
	},
}

func init() {

	ConfigCmd.Flags().String("client-id", "", "Contabo API Client ID")
	ConfigCmd.Flags().String("client-secret", "", "Contabo API Client Secret")
	ConfigCmd.Flags().String("api-user", "", "Contabo API User")
	ConfigCmd.Flags().String("api-password", "", "Contabo API Password")

	ConfigCmd.MarkFlagRequired("client-id")
	ConfigCmd.MarkFlagRequired("client-secret")
	ConfigCmd.MarkFlagRequired("api-user")
	ConfigCmd.MarkFlagRequired("api-password")
}
