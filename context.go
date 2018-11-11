package oksana

import (
	"net/http"
	"net/url"
)

// Context struct
type Context struct {
	Request  *http.Request
	Response *Response
	Params   *url.Values
}

// NewContext creates a new context
func NewContext() *Context {
	return &Context{}
}
