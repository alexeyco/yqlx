package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command.
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate query builder",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("generate called")
	},
}

func init() {
	generateCmd.Flags().StringP("output", "o", "internal/yqlx", "Query builder output directory")

	rootCmd.AddCommand(generateCmd)
}
