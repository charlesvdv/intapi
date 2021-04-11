package schemas

// TODO: would schemas.Generic wouldn't be better than schemas.Schema?

// A Schema declares a generic interface that any schema type must
// adhere to.
type Schema interface {
	schemaImpl()
}
