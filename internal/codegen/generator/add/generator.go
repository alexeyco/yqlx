package add

import (
	"fmt"
	"path"

	"github.com/gobeam/stringy"
	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/templatex"
)

// Generator describes a new table generator.
type Generator struct{}

// New returns a new Generator.
func New() *Generator {
	return &Generator{}
}

// Generate generates a new table.
func (g *Generator) Generate(fs afero.Fs, output, tablePath string) error {
	packageName := path.Base(output)
	tableBaseName := path.Base(tablePath)
	tableDirName := path.Dir(tablePath)

	tableGoName := fmt.Sprintf("%s.go", stringy.New(tableBaseName).SnakeCase().ToLower())
	tableTypeName := stringy.New(tableBaseName).CamelCase()

	return templatex.CompileWrite(fs, path.Join(output, tableGoName), tpl,
		templatex.WithoutHeader,
		templatex.Data(map[string]any{
			"Package": packageName,
			"Type":    tableTypeName,
			"Path":    fmt.Sprintf("%q", tableDirName),
			"Table":   fmt.Sprintf("%q", tableBaseName),
		}))
}
