package cmd

import (
	"folly/src/lib/generators"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs the tests",
	Long:  `Runs the tests.`,
	Run: func(cmd *cobra.Command, args []string) {
		contents := generators.GenerateModel("Aguacate")

		println(contents)

		os.WriteFile("src/app/models/test.go", []byte(contents), 0644)

	},
}
