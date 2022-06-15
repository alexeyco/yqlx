package schema

import "github.com/alexeyco/yqlx/schema/types"

// Column describes YDB table column.
type Column struct {
	Kind       types.Type
	Name       string
	IsOptional bool
}

// Optional to make a column optional.
func (c *Column) Optional() *Column {
	c.IsOptional = true

	return c
}

// Bool returns a new bool column.
func Bool(name string) *Column {
	return &Column{
		Name: name,
	}
}

// String returns a new string column.
func String(name string) *Column {
	return &Column{
		Name: name,
	}
}
