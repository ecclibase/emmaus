package oksana

import "encoding/json"

// Oksana struct holds router and context for framework
type Oksana struct {
	Config     Config
	Context    *Context
	Middleware []MiddlewareHandler
	Router     *Router
}

// Handler basic function to router handlers
type Handler func(*Context) error

// MiddlewareHandler defines a function to process middleware
type MiddlewareHandler func(Handler) Handler

// NotFoundHandler default 404 handler for not found routes
func NotFoundHandler(context *Context) (err error) {
	b, _ := json.Marshal("Not Found")

	context.Response.Header().Set("Content-Type", "application/json")
	context.Response.WriteHeader(404)
	_, err = context.Response.Write([]byte(b))

	return
}

// New creates a new service
func New() (oksana *Oksana) {
	return &Oksana{
		Context: NewContext(),
		Router:  NewRouter(),
	}
}

// GetContext returns current Context
func (oksana *Oksana) GetContext() *Context {
	return oksana.Context
}

// MiddlewareHandler adds a middlware handler
func (oksana *Oksana) MiddlewareHandler(middleware ...MiddlewareHandler) {
	for _, handler := range middleware {
		oksana.Middleware = append(oksana.Middleware, handler)
	}
}

// DELETE adds a HTTP Delete route to router
func (oksana *Oksana) DELETE(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("DELETE", endpoint, handler, middleware)
}

// GET adds a HTTP Get route to router
func (oksana *Oksana) GET(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("GET", endpoint, handler, middleware)
}

func (oksana *Oksana) addroute(action string, endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add(action, endpoint, handler, middleware)
}
