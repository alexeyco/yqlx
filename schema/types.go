package schema

// Type describes type interface.
type Type interface {
	isType()
}

type primitiveType uint8

func (primitiveType) isType() {}

const (
	// TypeUnknown unknown type.
	TypeUnknown primitiveType = iota
	// TypeBool bool type.
	TypeBool
	// TypeInt8 int8 type.
	TypeInt8
	// TypeUint8 uint8 type.
	TypeUint8
	// TypeInt16 int16 type.
	TypeInt16
	// TypeUint16 uint16 type.
	TypeUint16
	// TypeInt32 int32 type.
	TypeInt32
	// TypeUint32 uint32 type.
	TypeUint32
	// TypeInt64 int64 type.
	TypeInt64
	// TypeUint64 uint64 type.
	TypeUint64
	// TypeFloat float type.
	TypeFloat
	// TypeDouble double type.
	TypeDouble
	// TypeDate date type.
	TypeDate
	// TypeDatetime datetime type.
	TypeDatetime
	// TypeTimestamp timestamp type.
	TypeTimestamp
	// TypeInterval interval type.
	TypeInterval
	// TypeTzDate tzdate type.
	TypeTzDate
	// TypeTzDatetime tzdatetime type.
	TypeTzDatetime
	// TypeTzTimestamp tztimestamp type.
	TypeTzTimestamp
	// TypeString string type.
	TypeString
	// TypeUTF8 utf8 type.
	TypeUTF8
	// TypeYSON yson type.
	TypeYSON
	// TypeJSON json type.
	TypeJSON
	// TypeUUID uuid type.
	TypeUUID
	// TypeJSONDocument jsondocument type.
	TypeJSONDocument
	// TypeDyNumber dynumber type.
	TypeDyNumber
)
