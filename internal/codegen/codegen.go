package codegen

import (
	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/generator/add"
)

// Codegen code generator.
type Codegen struct {
	fs        afero.Fs
	validator Validator
	add       AddGenerator
}

// New returns new Codegen.
func New(opts ...Option) *Codegen {
	fs := afero.NewOsFs()

	c := &Codegen{
		fs:        fs,
		validator: NewDefaultValidator(),
		add:       add.New(),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Add generates a new table.
func (c *Codegen) Add(output, tablePath string) error {
	if err := c.validator.ValidateAdd(c.fs, output, tablePath); err != nil {
		return err
	}

	err := c.add.Generate(c.fs, output, tablePath)

	return err
}
