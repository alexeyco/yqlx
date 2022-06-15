package codegen

import "errors"

var (
	// ErrInternal internal error.
	ErrInternal = errors.New("something went wrong")
	// ErrOutput returns when a problem with output directory.
	ErrOutput = errors.New("output directory")
	// ErrTable returns when a problem with table path.
	ErrTable = errors.New("table path")
)
