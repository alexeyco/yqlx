package utils

import (
	"fmt"
	"path/filepath"

	"github.com/gobeam/stringy"
)

// PackageName returns package name by path.
func PackageName(s string) (string, error) {
	abs, err := filepath.Abs(s)
	if err != nil {
		return "", err
	}

	packageName := filepath.Base(abs)
	if packageName == "" || packageName == string(filepath.Separator) {
		packageName = "schema"
	}

	return packageName, nil
}

// TypeName returns type name by table name.
func TypeName(s string) string {
	return stringy.New(s).CamelCase()
}

// TableName returns table name by table name (he-he).
func TableName(s string) string {
	return stringy.New(s).SnakeCase().ToLower()
}

// FileName returns file name by table.
func FileName(s string) string {
	return fmt.Sprintf("%s.go", TableName(s))
}
