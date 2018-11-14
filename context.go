package oksana

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Context struct
type Context struct {
	Request  *http.Request
	Response *Response
	Params   url.Values
}

// NewContext creates a new context struct
func NewContext() *Context {
	return &Context{}
}

// SetRequest add request to context field
func (context *Context) SetRequest(request *http.Request) {
	context.Request = request
	context.Params = request.URL.Query()
}

// SetResponse add response to context field
func (context *Context) SetResponse(writer http.ResponseWriter) {
	context.Response = NewResponse(writer)
}

// Code writes header with HTTP code
func (context *Context) Code(code int) (err error) {
	context.Response.WriteHeader(code)
	return nil
}

// JSON returns response as serialized JSON
func (context *Context) JSON(code int, i interface{}) (err error) {
	b, err := json.Marshal(i)

	if err != nil {
		context.HTTPError(500, err.Error())
		return
	}

	context.SetHeader("Content-Type", "application/json")
	context.Response.WriteHeader(code)
	_, err = context.Response.Write([]byte(b))
	return
}

// String returns response as a string
func (context *Context) String(code int, s string) (err error) {
	context.SetHeader("Content-Type", "text/html;charset=utf-8")
	context.Response.WriteHeader(code)
	_, err = context.Response.Write([]byte(s))
	return
}

// GetParam return specified paramater
func (context *Context) GetParam(param string) string {
	return context.Params.Get(param)
}

// HasParam checks if param is set
func (context *Context) HasParam(param string) bool {
	if context.GetParam(param) != "" {
		return true
	}

	return false
}

// HTTPError returns a text/html error with requested code
func (context *Context) HTTPError(code int, message string) (err error) {
	context.Response.Header().Set("Content-Type", "text/html;charset=utf-8")
	context.Response.WriteHeader(code)
	_, err = context.Response.Write([]byte(message))
	return
}

// GetHeader returns specified header
func (context *Context) GetHeader(header string) string {
	return context.Request.Header.Get(header)
}

// SetHeader adds header to response
func (context *Context) SetHeader(k string, v string) {
	context.Response.Header().Set(k, v)
}

// Redirect returns a HTTP code
func (context *Context) Redirect(code int, uri string) error {
	http.Redirect(context.Response.Writer, context.Request, uri, code)
	return nil
}
