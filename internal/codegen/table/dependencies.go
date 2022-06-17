package table

import "github.com/spf13/afero"

// Validator describes validator interface.
type Validator interface {
	Validate(afero.Fs, string, string) error
}

// Writer describes writer interface.
type Writer interface {
	Write(afero.Fs, string, string) error
}
