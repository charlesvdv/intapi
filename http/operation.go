package http

import (
	"strings"

	"github.com/charlesvdv/servapigen/schemas"
)

type OperationDefinition struct {
	path            string
	method          string
	pathExpressions []PathExpressionDefinition
	description     string
	request         *RequestDefinition
	responses       []ResponseDefinition
}

func (op *OperationDefinition) Description(description string) *OperationDefinition {
	op.description = strings.TrimSpace(description)
	return op
}

func (op *OperationDefinition) Request(request *RequestDefinition) *OperationDefinition {
	op.request = request
	return op
}

func (op *OperationDefinition) Responses(responses ...*ResponseDefinition) *OperationDefinition {
	for _, resp := range responses {
		op.responses = append(op.responses, *resp)
	}
	return op
}

func Request() *RequestDefinition {
	return &RequestDefinition{}
}

type RequestDefinition struct {
	queryStrings []QueryStringDefinition
}

func (req *RequestDefinition) QueryStrings(queryStrings ...*QueryStringDefinition) *RequestDefinition {
	for _, q := range queryStrings {
		if q == nil {
			continue
		}
		req.queryStrings = append(req.queryStrings, *q)
	}
	return req
}

func BodyContent() *BodyContentDefinition {
	return &BodyContentDefinition{}
}

func JsonContent(schema schemas.Schema) *BodyContentDefinition {
	return &BodyContentDefinition{
		mime:   "application/json",
		schema: schema,
	}
}

type BodyContentDefinition struct {
	mime   string
	schema schemas.Schema
}

func QueryString(name string, schema schemas.Schema) *QueryStringDefinition {
	return &QueryStringDefinition{
		name:   name,
		schema: schema,
	}
}

type QueryStringDefinition struct {
	name        string
	required    bool
	schema      schemas.Schema
	description string
}

func (qs *QueryStringDefinition) Description(desc string) *QueryStringDefinition {
	qs.description = desc
	return qs
}

func PathExpression(name string, schema schemas.Schema) *PathExpressionDefinition {
	return &PathExpressionDefinition{
		name:   name,
		schema: schema,
	}
}

type PathExpressionDefinition struct {
	name        string
	schema      schemas.Schema
	description string
}

func (pe *PathExpressionDefinition) Description(desc string) *PathExpressionDefinition {
	pe.description = desc
	return pe
}

func Header(name string, schema schemas.Schema) *HeaderDefinition {
	return &HeaderDefinition{
		name:   name,
		schema: schema,
	}
}

type HeaderDefinition struct {
	name        string
	schema      schemas.Schema
	description string
}

func (h *HeaderDefinition) Description(desc string) *HeaderDefinition {
	h.description = desc
	return h
}
