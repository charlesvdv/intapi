package http

import nethttp "net/http"

const (
	// Invalid HTTP status that will act as the catch-all definition is
	// no other response definition is matching.
	StatusCatchAll = 0
)

func Response(status int16) *ResponseDefinition {
	return &ResponseDefinition{
		status: status,
	}
}

// TODO: we should probably generate those automatically and
// create one for every http status code

func ResponseOk() *ResponseDefinition {
	return Response(nethttp.StatusOK)
}

func ResponseCreated() *ResponseDefinition {
	return Response(nethttp.StatusCreated)
}

func ResponseCatchAll() *ResponseDefinition {
	return Response(StatusCatchAll)
}

type ResponseDefinition struct {
	status      int16
	description string
	contents    []BodyContentDefinition
	headers     []HeaderDefinition
}

func (resp *ResponseDefinition) Description(desc string) *ResponseDefinition {
	resp.description = desc
	return resp
}

func (resp *ResponseDefinition) Contents(contents ...*BodyContentDefinition) *ResponseDefinition {
	for _, c := range contents {
		if c == nil {
			continue
		}
		resp.contents = append(resp.contents, *c)
	}
	return resp
}

func (resp *ResponseDefinition) Headers(headers ...*HeaderDefinition) *ResponseDefinition {
	for _, h := range headers {
		if h == nil {
			continue
		}
		resp.headers = append(resp.headers, *h)
	}
	return resp
}
