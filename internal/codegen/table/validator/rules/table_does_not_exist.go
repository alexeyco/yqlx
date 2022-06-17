package rules

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/utils"
	"github.com/alexeyco/yqlx/internal/codegen/table/validator/errors"
)

// TableDoesNotExist checks if table already exists.
func TableDoesNotExist(fs afero.Fs, tableName, directory string) error {
	fileName := path.Join(directory, utils.FileName(tableName))

	_, err := fs.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	return fmt.Errorf("table file %q %w", fileName, errors.ErrTableAlreadyExists)
}
