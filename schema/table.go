package schema

// Table describes YDB table.
type Table interface {
	Path() string
	Name() string
	Columns() []*Column
}
