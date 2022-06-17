package table

import (
	"os"

	"github.com/spf13/afero"

	"github.com/alexeyco/yqlx/internal/codegen/table/validator"
	"github.com/alexeyco/yqlx/internal/codegen/table/writer"
)

// Generator describes table generator.
type Generator struct {
	tableName string
	fs        afero.Fs
	validator Validator
	writer    Writer
	logger    Logger
}

// Generate returns new Generator.
func Generate(tableName string, opts ...Option) *Generator {
	g := &Generator{
		tableName: tableName,
		fs:        afero.NewOsFs(),
		validator: validator.New(),
		writer:    writer.New(),
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// To generates table into specified directory.
func (g *Generator) To(directory string) error {
	if err := g.fs.MkdirAll(directory, os.ModePerm); err != nil {
		return err
	}

	if err := g.validator.Validate(g.fs, g.tableName, directory); err != nil {
		return err
	}

	return g.writer.Write(g.fs, g.tableName, directory, func(fileName string) {
		if g.logger != nil {
			g.logger.Infof("Created file %q", fileName)
		}
	})
}
