package schema

// Column describes YDB table column.
type Column struct {
	IsOptional bool
	Type       Type
	Name       string
}

// Optional to make a column optional.
func (c *Column) Optional() *Column {
	c.IsOptional = true

	return c
}

// Bool returns bool Column.
func Bool(name string) *Column { return &Column{Type: TypeBool, Name: name} }

// Int8 returns int8 Column.
func Int8(name string) *Column { return &Column{Type: TypeInt8, Name: name} }

// Int16 returns int16 Column.
func Int16(name string) *Column { return &Column{Type: TypeInt16, Name: name} }

// Int32 returns int32 Column.
func Int32(name string) *Column { return &Column{Type: TypeInt32, Name: name} }

// Int64 returns int64 Column.
func Int64(name string) *Column { return &Column{Type: TypeInt64, Name: name} }

// Uint8 returns uint8 Column.
func Uint8(name string) *Column { return &Column{Type: TypeUint8, Name: name} }

// Uint16 returns uint16 Column.
func Uint16(name string) *Column { return &Column{Type: TypeUint16, Name: name} }

// Uint32 returns uint32 Column.
func Uint32(name string) *Column { return &Column{Type: TypeUint32, Name: name} }

// Uint64 returns uint64 Column.
func Uint64(name string) *Column { return &Column{Type: TypeUint64, Name: name} }

// Float returns float Column.
func Float(name string) *Column { return &Column{Type: TypeFloat, Name: name} }

// Double returns double Column.
func Double(name string) *Column { return &Column{Type: TypeDouble, Name: name} }

// Date returns date Column.
func Date(name string) *Column { return &Column{Type: TypeDate, Name: name} }

// Datetime returns datetime Column.
func Datetime(name string) *Column { return &Column{Type: TypeDatetime, Name: name} }

// Timestamp returns timestamp Column.
func Timestamp(name string) *Column { return &Column{Type: TypeTimestamp, Name: name} }

// Interval returns interval Column.
func Interval(name string) *Column { return &Column{Type: TypeInterval, Name: name} }

// TzDate returns tzdate Column.
func TzDate(name string) *Column { return &Column{Type: TypeTzDate, Name: name} }

// TzDatetime returns tzdatetime Column.
func TzDatetime(name string) *Column { return &Column{Type: TypeTzDatetime, Name: name} }

// TzTimestamp returns tztimestamp Column.
func TzTimestamp(name string) *Column { return &Column{Type: TypeTzTimestamp, Name: name} }

// String returns string Column.
func String(name string) *Column { return &Column{Type: TypeString, Name: name} }

// UTF8 returns utf8 Column.
func UTF8(name string) *Column { return &Column{Type: TypeUTF8, Name: name} }

// YSON returns yson Column.
func YSON(name string) *Column { return &Column{Type: TypeYSON, Name: name} }

// JSON returns json Column.
func JSON(name string) *Column { return &Column{Type: TypeJSON, Name: name} }

// UUID returns uuid Column.
func UUID(name string) *Column { return &Column{Type: TypeUUID, Name: name} }

// JSONDocument returns jsondocument Column.
func JSONDocument(name string) *Column { return &Column{Type: TypeJSONDocument, Name: name} }

// DyNumber returns dynumber Column.
func DyNumber(name string) *Column { return &Column{Type: TypeDyNumber, Name: name} }
