package cmd

import (
	"fmt"

	"folly/src/database"
	"folly/src/database/factories"

	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Coneecting to database...")
		if _, err := database.Connect(); err != nil {
			panic(err)
		}
		fmt.Println("Connected successfully!")

		fmt.Println("Seeding the database...")

		// Create author role
		authorRole, err := factories.RoleFactory.CreateOneWith(map[string]interface{}{
			"Name": "author",
		})
		if err != nil {
			panic(err)
		}

		// Create 100 authors
		authors, err := factories.UserFactory.CreateManyWith(100, map[string]interface{}{
			"Roles": []interface{}{authorRole},
		})
		if err != nil {
			panic(err)
		}

		// Create 100 posts for each author
		for _, author := range authors {
			derefAutor := **author
			_, err := factories.PostFactory.CreateManyWith(100, map[string]interface{}{
				"AuthorID": derefAutor.ID,
			})
			if err != nil {
				panic(err)
			}

		}

		fmt.Println("Database seeded successfully!")
	},
}
