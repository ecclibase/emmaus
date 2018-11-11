package oksana

// Route holds all information about a defined route
type Route struct {
	Handler    Handler
	Method     string
	Path       string
	Middleware []MiddlewareHandler
}

// Router holds all defined routes
type Router struct {
	Routes map[string]Route
}

// New creates a new router
func (router *Router) New() {
	return &Router{
		Routes: make(map[string]Route),
	}
}

// Add will add a new route to the Router.Routes map
func (router *Router) Add(method string, path string, handler Handler, middleware []MiddlewareHandler) {
	route := Route{
		Handler: handler,
		Method:  method,
		Path:    path,
	}

	// add middleware handler(s)
	for _, v := range middleware {
		route.Middleware = append(route.Middleware, v)
	}

	router.Routes[method+path] = route
}

// GetRoute will search routes and return Route
func (router *Router) GetRoute(context *Context) Route {
	var route Route

	path := context.Request.URL.Path
	method := context.Request.Method

	if val, ok := router.Routes[method+path]; ok {
		route = router.Routes[method+path]
	}

	return route
}
