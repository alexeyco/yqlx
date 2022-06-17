package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexeyco/yqlx/internal/codegen/table/utils"
)

var packageNameData = [...]struct {
	name     string
	source   string
	expected string
}{
	{
		name:     "Absolute",
		source:   "/path/to/directory",
		expected: "directory",
	},
	{
		name:     "Root",
		source:   "/",
		expected: "schema",
	},
}

func TestPackageName(t *testing.T) {
	t.Parallel()

	for _, datum := range packageNameData {
		datum := datum

		t.Run(datum.name, func(t *testing.T) {
			t.Parallel()

			actual, err := utils.PackageName(datum.source)

			assert.NoError(t, err)
			assert.Equal(t, datum.expected, actual)
		})
	}
}

var typeNameData = [...]struct {
	source   string
	expected string
}{
	{
		source:   "table_name",
		expected: "TableName",
	},
	{
		source:   "TableName",
		expected: "TableName",
	},
}

func TestTypeName(t *testing.T) {
	t.Parallel()

	for _, datum := range typeNameData {
		datum := datum

		t.Run(datum.source, func(t *testing.T) {
			t.Parallel()

			actual := utils.TypeName(datum.source)

			assert.Equal(t, datum.expected, actual)
		})
	}
}

var tableNameData = [...]struct {
	source   string
	expected string
}{
	{
		source:   "table_name",
		expected: "table_name",
	},
	{
		source:   "TableName",
		expected: "table_name",
	},
}

func TestTableName(t *testing.T) {
	t.Parallel()

	for _, datum := range tableNameData {
		datum := datum

		t.Run(datum.source, func(t *testing.T) {
			t.Parallel()

			actual := utils.TableName(datum.source)

			assert.Equal(t, datum.expected, actual)
		})
	}
}

var fileNameData = [...]struct {
	source   string
	expected string
}{
	{
		source:   "table_name",
		expected: "table_name.go",
	},
	{
		source:   "TableName",
		expected: "table_name.go",
	},
}

func TestFileName(t *testing.T) {
	t.Parallel()

	for _, datum := range fileNameData {
		datum := datum

		t.Run(datum.source, func(t *testing.T) {
			t.Parallel()

			actual := utils.FileName(datum.source)

			assert.Equal(t, datum.expected, actual)
		})
	}
}
