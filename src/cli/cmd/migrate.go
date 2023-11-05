package cmd

import (
	"fmt"
	"folly/src/database"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates the database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coneecting to database...")
		if _, err := database.Connect(); err != nil {
			panic(err)
		}
		fmt.Println("Connected successfully!")
		fmt.Println("Migrating the database...")
		if err := database.Migrate(); err != nil {
			panic(err)
		}
		fmt.Println("Database migrated successfully!")
	},
}
