package oksana

// Group holds information about the route group
type Group struct {
	Middleware []MiddlewareHandler
	Prefix     string
	Router     Router
}

// DELETE adds a HTTP Get method to the group
func (group *Group) DELETE(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("DELETE", path, handler, middleware)
}

// GET adds a HTTP Get method to the group
func (group *Group) GET(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("GET", path, handler, middleware)
}

// OPTIONS adds a HTTP Get method to the group
func (group *Group) OPTIONS(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("OPTIONS", path, handler, middleware)
}

// PATCH adds a HTTP Get method to the group
func (group *Group) PATCH(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("PATCH", path, handler, middleware)
}

// POST adds a HTTP Get method to the group
func (group *Group) POST(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("POST", path, handler, middleware)
}

// PUT adds a HTTP Get method to the group
func (group *Group) PUT(path string, handler Handler, middleware ...MiddlewareHandler) {
	group.add("PUT", path, handler, middleware)
}

func (group *Group) add(method string, path string, handler Handler, middleware []MiddlewareHandler) {
	path = group.Prefix + path
	group.Router.Add(method, path, handler, middleware)
}
