package petstore

import "github.com/charlesvdv/intapi/schemas"

var petSchema = schemas.Object().
	Element("id", schemas.Integer()).
	Element("name", schemas.String()).
	OptionalElement("tag", schemas.String())

var errorSchema = schemas.Object().
	Element("code", schemas.Integer()).
	Element("message", schemas.String())
