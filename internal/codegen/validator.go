package codegen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/gobeam/stringy"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/afero"
)

// Validator describes Codegen validator.
type Validator interface {
	ValidateAdd(afero.Fs, string, string) error
}

var nameRegexp = regexp.MustCompile(`^/[\w\d-/_]+$`)

type validator func() error

// DefaultValidator describes a default Codegen validator.
type DefaultValidator struct{}

// NewDefaultValidator returns new DefaultValidator.
func NewDefaultValidator() *DefaultValidator {
	return &DefaultValidator{}
}

// ValidateAdd validates data before Codegen generates a new table.
func (v *DefaultValidator) ValidateAdd(fs afero.Fs, output, tablePath string) (res error) {
	validators := []validator{
		v.outputIsNotFile(fs, output),
		v.pathIsValid(tablePath),
		v.tableDoesNotExist(fs, output, tablePath),
	}

	for _, validate := range validators {
		if err := validate(); err != nil {
			res = multierror.Append(res, err)
		}
	}

	return
}

func (*DefaultValidator) outputIsNotFile(fs afero.Fs, output string) validator {
	return func() error {
		fi, err := fs.Stat(output)
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}

			return fmt.Errorf("%w: %s", ErrInternal, err)
		}

		if fi.IsDir() {
			return nil
		}

		return fmt.Errorf("%w shouldn't be a file", ErrOutput)
	}
}

func (*DefaultValidator) pathIsValid(tablePath string) validator {
	return func() error {
		if nameRegexp.MatchString(tablePath) {
			return nil
		}

		return fmt.Errorf("%w should contain only letters", ErrTable)
	}
}

func (*DefaultValidator) tableDoesNotExist(fs afero.Fs, output, tablePath string) validator {
	tableName := stringy.New(filepath.Base(tablePath)).SnakeCase().ToLower()
	fileName := fmt.Sprintf("%s.go", tableName)

	return func() error {
		_, err := fs.Stat(path.Join(output, fileName))
		if err != nil {
			if os.IsNotExist(err) {
				return nil
			}

			return fmt.Errorf("%w: %s", ErrInternal, err)
		}

		return fmt.Errorf("%w already exists", ErrTable)
	}
}
