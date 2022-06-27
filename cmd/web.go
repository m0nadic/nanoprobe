package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nanoprobe/internal/app/server"
	"nanoprobe/internal/app/utils"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Starts the web interface",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called with configs", cfgFile)
		cfg, err := utils.LoadConfig(cfgFile)

		if err != nil {
			utils.ErrorExit(1, err)
		}

		err = server.StartWebServer(cfg)
		if err != nil {
			utils.ErrorExit(1, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
