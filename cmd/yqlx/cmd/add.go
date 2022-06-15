package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/alexeyco/yqlx/internal/codegen"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:     "add path/to/table",
	Short:   "Add a table",
	Example: "yqlx add path/to/table -o internal/yqlx/schema",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := codegen.New().Add(cmd.Flag("output").Value.String(), args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	addCmd.Flags().StringP("output", "o", "internal/yqlx/schema", "Schema directory")

	rootCmd.AddCommand(addCmd)
}
