package table

import "github.com/spf13/afero"

// Option describes Generator option.
type Option func(*Generator)

// WithFs to set custom fs.
func WithFs(fs afero.Fs) Option {
	return func(g *Generator) {
		g.fs = fs
	}
}

// WithValidator to set custom Validator.
func WithValidator(v Validator) Option {
	return func(g *Generator) {
		g.validator = v
	}
}

// WithWriter to set custom Writer.
func WithWriter(w Writer) Option {
	return func(g *Generator) {
		g.writer = w
	}
}

// WithLogger to set custom Logger.
func WithLogger(l Logger) Option {
	return func(g *Generator) {
		g.logger = l
	}
}
