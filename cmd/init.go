package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nanoprobe/internal/app/configs"
	"nanoprobe/internal/app/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a basic setup for nanoprobe.",
	Long:  `Generates a sample config file and other dependencies`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("server", &configs.ServerConfig{
			Host: "localhost",
			Port: 9876,
		})
		viper.Set("database", &configs.DatabaseConfig{DBFile: "./nanoprobe.yml"})
		err := viper.WriteConfigAs("nanoprobe.yaml")
		if err != nil {
			utils.ErrorExit(1, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
