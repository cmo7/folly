package cmd

import (
	"folly/src/database"
	"folly/src/database/seeders"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := database.Connect(); err != nil {
			panic(err)
		}
		if err := seeders.Seed(); err != nil {
			panic(err)
		}
	},
}
