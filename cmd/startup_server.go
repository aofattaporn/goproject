package cmd

import (
	"fmt"

	"github.com/goproject/configs"
	"github.com/spf13/cobra"
)

// command startUpCmd usage
var startUpCmd = &cobra.Command{
	Use:   "startup",
	Short: "Startup Server",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Loading config
		_, err := configs.LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("load config error: %+v", err)
		}

		// excute run on
		fmt.Println("Test Hello world")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startUpCmd)
}
