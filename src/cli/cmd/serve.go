package cmd

import (
	"fmt"
	"folly/src/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Folly server",
	Long:  `Starts the Folly server, usefull for development.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Folly server...")

		app.Serve()
	},
}
