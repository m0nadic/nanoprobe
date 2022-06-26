package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nanoprobe/internal/app/utils"
	"os"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Starts the web interface",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called with configs", cfgFile)
		cfg, err := utils.LoadConfig(cfgFile)

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
		}

		fmt.Println(cfg)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
