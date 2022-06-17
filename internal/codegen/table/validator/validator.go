package validator

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/validator/rules"
)

// Rule describes validation rule.
type Rule func(afero.Fs, string, string) error

// Validator describes validator.
type Validator []Rule

// New returns new Validator.
func New() Validator {
	return Validator{
		rules.TableNameIsCorrect,
		rules.DirectoryExists,
		rules.TableDoesNotExist,
	}
}

// Validate validates.
func (v Validator) Validate(fs afero.Fs, tableName, directory string) (res error) {
	for _, rule := range v {
		if err := rule(fs, tableName, directory); err != nil {
			res = multierror.Append(res, err)
		}
	}

	return
}
