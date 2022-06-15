package codegen

import "github.com/spf13/afero"

// AddGenerator describes a generator that generates a new table.
type AddGenerator interface {
	Generate(afero.Fs, string, string) error
}
