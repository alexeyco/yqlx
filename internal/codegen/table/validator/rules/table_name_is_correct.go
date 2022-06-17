package rules

import (
	"fmt"
	"regexp"

	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/validator/errors"
)

var tableNameRegexp = regexp.MustCompile(`^[\w\d-_]+$`)

// TableNameIsCorrect checks table name.
func TableNameIsCorrect(fs afero.Fs, tableName, _ string) error {
	if tableNameRegexp.MatchString(tableName) {
		return nil
	}

	return fmt.Errorf("table name is %w", errors.ErrIncorrectTableName)
}
