package add

import "text/template"

var tpl = template.Must(template.New("tpl").Parse(`package {{ .Package }}

import "github.com/alexeyco/yqlx/schema"

type {{ .Type }} struct {
	schema.Table
}

func ({{ .Type }}) Path() string {
	return {{ .Path }}
}

func ({{ .Type }}) Name() string {
	return {{ .Table }}
}

func ({{ .Type }}) Columns() []*schema.Column {
	return []*schema.Column{
		schema.String("id"),
		schema.Bool("is_active"),
	}
}
`))
