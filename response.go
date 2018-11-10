package oksana

import "net/http"

// Response standard Oksana response struct
type Response struct {
	Committed bool
	Size      int64
	Status    int
	Writer    http.ResponseWriter
}
