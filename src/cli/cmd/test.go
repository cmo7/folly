package cmd

import (
	"fmt"

	"github.com/iancoleman/strcase"
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
		fields := []string{"first-name", "first_name", "UpdatedAt", "DeletedAt"}
		for _, field := range fields {
			fmt.Println(field, "=>", strcase.ToCamel(field))
			fmt.Println(field, "=>", strcase.ToKebab(field))
		}
	},
}
