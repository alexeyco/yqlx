package codegen_test

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen"
)

// nolint:funlen
func TestDefaultValidator_ValidateAdd(t *testing.T) {
	t.Parallel()

	const (
		output    = "internal/yqlx/schema"
		tablePath = "/path/to/table"
	)

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewMemMapFs()

		err := codegen.NewDefaultValidator().
			ValidateAdd(fs, output, tablePath)

		assert.NoError(t, err)
	})

	t.Run("DirectoryExists", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewMemMapFs()

		err := fs.MkdirAll(output, os.ModePerm)
		assert.NoError(t, err)

		err = codegen.NewDefaultValidator().
			ValidateAdd(fs, output, tablePath)

		assert.NoError(t, err)
	})

	t.Run("OutputFileExists", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewMemMapFs()

		err := fs.MkdirAll(filepath.Dir(output), os.ModePerm)
		assert.NoError(t, err)

		f, err := fs.Create(output)
		assert.NoError(t, err)

		err = f.Close()
		assert.NoError(t, err)

		err = codegen.NewDefaultValidator().
			ValidateAdd(fs, output, tablePath)

		assert.ErrorIs(t, err, codegen.ErrOutput)
	})

	t.Run("IncorrectTable", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewMemMapFs()

		err := codegen.NewDefaultValidator().
			ValidateAdd(fs, output, "foo 123")

		assert.ErrorIs(t, err, codegen.ErrTable)
	})

	t.Run("TableExists", func(t *testing.T) {
		t.Parallel()

		fs := afero.NewMemMapFs()

		err := fs.MkdirAll(output, os.ModePerm)
		assert.NoError(t, err)

		f, err := fs.Create(path.Join(output, "table.go"))
		assert.NoError(t, err)

		err = f.Close()
		assert.NoError(t, err)

		err = codegen.NewDefaultValidator().
			ValidateAdd(fs, output, tablePath)

		assert.ErrorIs(t, err, codegen.ErrTable)
	})
}
