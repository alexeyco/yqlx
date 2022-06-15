package codegen

import "github.com/spf13/afero"

// Option describes Codegen option.
type Option func(*Codegen)

// WithFs to pass a custom Fs to Codegen.
func WithFs(fs afero.Fs) Option {
	return func(c *Codegen) {
		c.fs = fs
	}
}

// WithValidator to pass a custom Validator to Codegen.
func WithValidator(v Validator) Option {
	return func(c *Codegen) {
		c.validator = v
	}
}

// WithAddGenerator to pass a custom AddGenerator to Codegen.
func WithAddGenerator(g AddGenerator) Option {
	return func(c *Codegen) {
		c.add = g
	}
}
