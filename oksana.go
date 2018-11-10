package oksana

import "encoding/json"

type Oksana struct {
	Config     Configuration
	Context    *CTX
	Middleware []MiddlewareHandler
	Router     *Router
}

// Handler basic function to router handlers
type Handler func(*CTX) error

// MiddlewareHandler defines a function to process middleware
type MiddlewareHandler func(Handler) Handler

// NotFoundHandler default 404 handler for not found routes
func NotFoundHandler(ctx *CTX) (err error) {
	b, _ := json.Marshal("Not Found")

	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.WriteHeader(404)
	_, err = ctx.Response.Write([]byte(b))

	return
}

// New creates a new service
func New() (oksana *Oksana) {
	return &Oksana{
		Context: new(Context),
		Router:  new(Router),
	}
}

// GetContext returns current Context
func (oksana *Oksana) GetContext() *Context {
	return oksana.Context
}

// Middleware adds a middlware handler
func (oksana *Oksana) Middleware(middleware ...MiddlewareHandler) {
	for i, handler := range middleware {
		oksana.Middleware = append(oksana.Middleware, middleware[i])
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
