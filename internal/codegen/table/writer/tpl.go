package writer

import "text/template"

var tpl = template.Must(template.New("tpl").Parse(`package {{ .Package }}

import "github.com/alexeyco/yqlx/schema"

var _ schema.Table = (*{{ .Type }})(nil)

// {{ .Type }} describes table {{ .Table | printf "%q" }}.
type {{ .Type }} struct {}

// Name returns table name.
func ({{ .Type }}) Name() string {
	return {{ .Table | printf "%q" }}
}

// Columns returns table columns.
func ({{ .Type }}) Columns() []*schema.Column {
	return []*schema.Column{
		// schema.String("string_column"),
		// schema.Bool("bool_column").Optional(),
		// Etc...
	}
}
`))
