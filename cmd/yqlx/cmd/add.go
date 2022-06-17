package cmd

import (
	"github.com/spf13/cobra"

	"github.com/alexeyco/yqlx/internal/codegen/table"
	"github.com/alexeyco/yqlx/internal/logx"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:     "add path/to/table",
	Short:   "Add a table",
	Example: "yqlx add path/to/table -o internal/yqlx/schema",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tableName := args[0]
		directory := cmd.Flag("output").Value.String()

		logger := logx.New()

		if err := table.Generate(tableName, table.WithLogger(logger)).To(directory); err != nil {
			logger.Fatal(err)
		}
	},
}

func init() {
	addCmd.Flags().StringP("output", "o", "internal/yqlx/schema", "Schema directory")

	rootCmd.AddCommand(addCmd)
}
