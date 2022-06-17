package writer

import (
	"path"

	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/utils"
	"github.com/alexeyco/yqlx/internal/templatex"
)

// Writer describes table writer.
type Writer struct{}

// New returns new Writer.
func New() *Writer {
	return &Writer{}
}

// Write generates a new table definition file and writes it to specified directory.
func (w *Writer) Write(fs afero.Fs, tableName, directory string) error {
	packageName, err := utils.PackageName(directory)
	if err != nil {
		return err
	}

	fileName := path.Join(directory, utils.FileName(tableName))

	return templatex.CompileWrite(fs, fileName, tpl,
		templatex.WithoutHeader,
		templatex.Data(map[string]any{
			"Package": packageName,
			"Type":    utils.TypeName(tableName),
			"Table":   utils.TableName(tableName),
		}))
}
