package schema

// Table describes YDB table.
type Table interface {
	Name() string
	Columns() []*Column
}
