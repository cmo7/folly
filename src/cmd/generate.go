package cmd

import (
	"folly/src/lib/generators"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates models, controllers, etc.",
	Long:  `Generates models, controllers, factories and repositories etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		modelName := args[0]
		generators.Generate(modelName)
	},
}
