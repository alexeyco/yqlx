package writer_test

import (
	"os"
	"path"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen/table/utils"
	"github.com/alexeyco/yqlx/internal/codegen/table/writer"
)

func TestWriter_Write(t *testing.T) {
	t.Parallel()

	const (
		tableName = "table_name"
		directory = "/path/to/directory"
	)

	fileName := path.Join(directory, utils.FileName(tableName))

	fsMock := afero.NewMemMapFs()

	err := fsMock.MkdirAll(directory, os.ModePerm)
	assert.NoError(t, err)

	err = writer.New().Write(fsMock, tableName, directory, func(s string) {
		assert.Equal(t, fileName, s)
	})
	assert.NoError(t, err)

	fi, err := fsMock.Stat(fileName)

	assert.NoError(t, err)
	assert.False(t, fi.IsDir())
}
