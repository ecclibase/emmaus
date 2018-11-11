package oksana

import "net/http"

// Response standard Oksana response struct
type Response struct {
	Committed bool
	Size      int64
	Status    int
	Writer    http.ResponseWriter
}

// NewResponse creates new Response struct
func NewResponse(w http.ResponseWriter) (r *Response) {
	return &Response{Writer: w}
}

// Write writs the bytes (message) to the client
func (response *Response) Write(b []byte) (n int, err error) {
	n, err = response.Writer.Write(b)
	return
}

// WriteHeader writes a header to the response writer
func (response *Response) WriteHeader(code int) {
	response.Writer.WriteHeader(code)
}

// Header returns the header information
func (response *Response) Header() http.Header {
	return response.Writer.Header()
}
