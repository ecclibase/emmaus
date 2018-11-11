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

// New creates a new context
func (ctx *Context) New() {
	return &Context{}
}
