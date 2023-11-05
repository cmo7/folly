package cmd

import (
	"folly/src/database"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(dropCmd)
}

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drops the database",
	Long:  `Drops the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Connect()
		database.Drop()
	},
}
