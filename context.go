package oksana

import "net/http"

// Context struct
type Context struct {
	Request  *http.Request
	Response *Response
	Params   map[string]string
}
