package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Folly",
	Long:  `Print the version number of Folly`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Folly CLI v0.0.1")
		fmt.Println("Config file:", viper.ConfigFileUsed())
	},
}
