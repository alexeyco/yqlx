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

	fsMock := afero.NewMemMapFs()

	err := fsMock.MkdirAll(directory, os.ModePerm)
	assert.NoError(t, err)

	err = writer.New().Write(fsMock, tableName, directory)
	assert.NoError(t, err)

	fi, err := fsMock.Stat(path.Join(directory, utils.FileName(tableName)))

	assert.NoError(t, err)
	assert.False(t, fi.IsDir())
}
