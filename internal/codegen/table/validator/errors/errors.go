package errors

import "errors"

var (
	// ErrIncorrectTableName when table name is incorrect.
	ErrIncorrectTableName = errors.New("incorrect")
	// ErrDirectoryDoesNotExist when directory doesn't exist.
	ErrDirectoryDoesNotExist = errors.New("doesn't exist")
	// ErrDirectoryShouldBeADirectory when output directory is a file.
	ErrDirectoryShouldBeADirectory = errors.New("should be a directory")
	// ErrTableAlreadyExists when table file already exists.
	ErrTableAlreadyExists = errors.New("already exists")
)
