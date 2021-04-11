package schemas

var (
	// Make sure the type declared in this file follows the `Type`
	// interface.
	_ Schema = StringSchema{}
	_ Schema = IntegerSchema{}
)

// String returns a StringSchema.
func String() StringSchema {
	return StringSchema{}
}

// A StringSchema defines a string type.
type StringSchema struct{}

func (s StringSchema) schemaImpl() {}

// Integer returns a IntegerSchema.
func Integer() IntegerSchema {
	return IntegerSchema{}
}

// An IntegerSchema defines a number type without a fractional component.
type IntegerSchema struct{}

func (i IntegerSchema) schemaImpl() {}
