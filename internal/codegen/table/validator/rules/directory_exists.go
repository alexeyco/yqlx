package rules

import (
	"fmt"
	"os"

	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/validator/errors"
)

// DirectoryExists checks if directory exists.
func DirectoryExists(fs afero.Fs, _, directory string) error {
	fi, err := fs.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("directory %q %w", directory, errors.ErrDirectoryDoesNotExist)
		}

		return err
	}

	if fi.IsDir() {
		return nil
	}

	return fmt.Errorf("path %q %w", directory, errors.ErrDirectoryShouldBeADirectory)
}
