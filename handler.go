package handler

import (
	"context"
	"net/http"
)

// Response of function call
type Response struct {

	// Body the body will be written back
	Body []byte

	// StatusCode needs to be populated with value such as http.StatusOK
	StatusCode int

	// Header is optional and contains any additional headers the function response should set
	Header http.Header
}

// Request of function call
type Request struct {
	Body        []byte
	Header      http.Header
	QueryString string
	Method      string
	Host        string

	ctx context.Context
}

// Context returns the request's context.
func (r *Request) Context() context.Context {
	return r.ctx
}

// WithContext changes the request's context.
func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r.ctx = ctx

	return r
}

// FunctionHandler used for a serverless Go method invocation
type FunctionHandler interface {
	Handle(req Request) (Response, error)
}

func init() {

}
