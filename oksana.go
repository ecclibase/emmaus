package oksana

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Oksana struct holds router and context for framework
type Oksana struct {
	Context    *Context
	Middleware []MiddlewareHandler
	Router     *Router
}

// Configuration struct holds values for configuring the application
type Configuration struct {
	Port string
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

// OPTIONS adds a HTTP Get route to router
func (oksana *Oksana) OPTIONS(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("OPTIONS", endpoint, handler, middleware)
}

// PATCH adds a HTTP Get route to router
func (oksana *Oksana) PATCH(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("PATCH", endpoint, handler, middleware)
}

// POST adds a HTTP Get route to router
func (oksana *Oksana) POST(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("POST", endpoint, handler, middleware)
}

// PUT adds a HTTP Get route to router
func (oksana *Oksana) PUT(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add("PUT", endpoint, handler, middleware)
}

// Group creates a route group with a common prefix
func (oksana *Oksana) Group(prefix string, middleware ...MiddlewareHandler) *Group {
	return &Group{
		Middleware: middleware,
		Prefix:     prefix,
		Router:     *oksana.Router,
	}
}

// Start initates the framework to start listening for requests
func (oksana *Oksana) Start(configuration ...Configuration) {
	var config Configuration
	if len(configuration) == 0 {
		config = Configuration{
			Port: "8080",
		}
	} else {
		config = configuration[0]
	}

	server := oksana.server(config)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %s", err)
	}
}

func (oksana *Oksana) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// handler
	var handler Handler

	// create context
	oksana.Context.SetRequest(request)
	oksana.Context.SetResponse(writer)

	if route, ok := oksana.Router.GetRoute(oksana.Context); ok {
		handler = func(*Context) error {
			handler := route.Handler

			if err := handler(oksana.Context); err != nil {
				panic(err)
			}

			return nil
		}

		handler(oksana.Context)
	} else {
		NotFoundHandler(oksana.Context)
	}

	return
}

func (oksana *Oksana) server(configuration Configuration) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", configuration.Port),
		Handler: oksana,
	}
}

func (oksana *Oksana) addroute(action string, endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	oksana.Router.Add(action, endpoint, handler, middleware)
}
