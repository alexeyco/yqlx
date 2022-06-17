package validator_test

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen/table/utils"
	"github.com/alexeyco/yqlx/internal/codegen/table/validator"
	"github.com/alexeyco/yqlx/internal/codegen/table/validator/errors"
)

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	const (
		tableName = "table_name"
		directory = "/path/to/directory"
	)

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(directory, os.ModePerm)
		assert.NoError(t, err)

		err = validator.New().Validate(fsMock, tableName, directory)

		assert.NoError(t, err)
	})

	t.Run("ErrTableAlreadyExists", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(directory, os.ModePerm)
		assert.NoError(t, err)

		f, err := fsMock.Create(path.Join(directory, utils.FileName(tableName)))
		assert.NoError(t, err)

		_ = f.Close()

		err = validator.New().Validate(fsMock, tableName, directory)

		assert.ErrorIs(t, err, errors.ErrTableAlreadyExists)
	})

	t.Run("ErrDirectoryShouldBeADirectory", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(path.Dir(directory), os.ModePerm)
		assert.NoError(t, err)

		f, err := fsMock.Create(directory)
		assert.NoError(t, err)

		_ = f.Close()

		err = validator.New().Validate(fsMock, tableName, directory)

		assert.ErrorIs(t, err, errors.ErrDirectoryShouldBeADirectory)
	})

	t.Run("ErrDirectoryDoesNotExist", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := validator.New().Validate(fsMock, tableName, directory)

		assert.ErrorIs(t, err, errors.ErrDirectoryDoesNotExist)
	})

	t.Run("ErrIncorrectTableName", func(t *testing.T) {
		t.Parallel()

		fsMock := afero.NewMemMapFs()

		err := fsMock.MkdirAll(directory, os.ModePerm)
		assert.NoError(t, err)

		err = validator.New().Validate(fsMock, "$foo@", directory)

		assert.ErrorIs(t, err, errors.ErrIncorrectTableName)
	})
}
