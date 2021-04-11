package http

import (
	nethttp "net/http"
)

func Get(path string, expressions ...*PathExpressionDefinition) *OperationDefinition {
	return newOperation(nethttp.MethodGet, path, expressions...)
}

func Post(path string, expressions ...*PathExpressionDefinition) *OperationDefinition {
	return newOperation(nethttp.MethodPost, path, expressions...)
}

func newOperation(method string, path string, expressions ...*PathExpressionDefinition) *OperationDefinition {
	var pathExpressions []PathExpressionDefinition
	for _, expr := range expressions {
		if expr == nil {
			continue
		}
		pathExpressions = append(pathExpressions, *expr)
	}

	return &OperationDefinition{
		method:          method,
		path:            path,
		pathExpressions: pathExpressions,
	}
}
