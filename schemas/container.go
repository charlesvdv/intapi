package schemas

func Object() *ObjectSchema {
	return &ObjectSchema{}
}

type ObjectSchema struct {
	elements []ObjectElement
}

func (obj *ObjectSchema) Element(name string, value Schema) *ObjectSchema {
	elem := ObjectElement{
		name:     name,
		value:    value,
		required: true,
	}
	obj.elements = append(obj.elements, elem)
	return obj
}

func (obj *ObjectSchema) OptionalElement(name string, value Schema) *ObjectSchema {
	elem := ObjectElement{
		name:     name,
		value:    value,
		required: false,
	}
	obj.elements = append(obj.elements, elem)
	return obj
}

func (obj *ObjectSchema) schemaImpl() {}

type ObjectElement struct {
	name     string
	value    Schema
	required bool
}

func Array(schema Schema) *ArraySchema {
	return &ArraySchema{
		schema: schema,
	}
}

type ArraySchema struct {
	schema Schema
}

func (arr *ArraySchema) schemaImpl() {}
