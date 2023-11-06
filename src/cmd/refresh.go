package cmd

import (
	"folly/src/database"
	"folly/src/database/seeders"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(refreshCmd)
}

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Drops, Migrates and Seeds the database",
	Long:  `Deletes every table, recreates them and seeds them using the main seeder.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := database.Connect(); err != nil {
			panic(err)
		}
		if err := database.Drop(); err != nil {
			panic(err)
		}
		if err := database.Migrate(); err != nil {
			panic(err)
		}
		if err := seeders.Seed(); err != nil {
			panic(err)
		}
	},
}
