package oksana

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedOksana struct {
	mock.Mock
}

func handler(context *Context) error {
	return nil
}

func middlware(handler Handler) Handler {
	return handler
}

func TestNewOksana(t *testing.T) {
	oksana := New()
	assert.NotNil(t, oksana)
}

func TestRouterIsOksanaRouter(t *testing.T) {
	oksana := New()
	assert.True(t, reflect.TypeOf(oksana.Router).String() == "*oksana.Router")
}

func TestAddMiddleware(t *testing.T) {
	oksana := New()

	middleware := func(handler Handler) Handler {
		return handler
	}

	oksana.MiddlewareHandler(middleware)

	assert.True(t, len(oksana.Middleware) == 1)
}

func TestCreateGroup(t *testing.T) {
	oksana := New()

	group := oksana.Group("/group")

	assert.True(t, reflect.TypeOf(group).String() == "*oksana.Group")
}

func TestGroupRouteAddedCorrectly(t *testing.T) {
	oksana := New()

	group := oksana.Group("/group")
	group.GET("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("GET", "/group/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "GET")
	assert.True(t, route.Path == "/group/test")
}

func TestAddDELETERoute(t *testing.T) {
	oksana := New()

	oksana.DELETE("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("DELETE", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "DELETE")
	assert.True(t, route.Path == "/test")
}

func TestAddGETRoute(t *testing.T) {
	oksana := New()

	oksana.GET("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("GET", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "GET")
	assert.True(t, route.Path == "/test")
}

func TestAddOPTIONSRoute(t *testing.T) {
	oksana := New()

	oksana.OPTIONS("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("OPTIONS", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "OPTIONS")
	assert.True(t, route.Path == "/test")
}

func TestAddPATCHRoute(t *testing.T) {
	oksana := New()

	oksana.PATCH("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("PATCH", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "PATCH")
	assert.True(t, route.Path == "/test")
}

func TestAddPOSTRoute(t *testing.T) {
	oksana := New()

	oksana.POST("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("POST", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "POST")
	assert.True(t, route.Path == "/test")
}

func TestAddPUTRoute(t *testing.T) {
	oksana := New()

	oksana.PUT("/test", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	route := oksana.Router.find("PUT", "/test")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
	assert.True(t, route.Method == "PUT")
	assert.True(t, route.Path == "/test")
}

func TestNotFoundHandler(t *testing.T) {
	oksana := New()

	oksana.GET("/", func(context *Context) error {
		return context.JSON(200, "This is a test")
	})

	r, _ := http.NewRequest("GET", "/blah", nil)
	w := httptest.NewRecorder()

	oksana.ServeHTTP(w, r)

	if assert.Equal(t, 404, w.Code) {
		assert.Equal(t, "\"Not Found\"", w.Body.String())
	}
}

func TestNewServerReturnsHTTPServer(t *testing.T) {
	config := Configuration{
		Port: "8080",
	}

	oksana := New()
	server := oksana.server(config)

	assert.True(t, reflect.TypeOf(server).String() == "*http.Server")
}

func TestNewServerRunsOnCorrectPort(t *testing.T) {
	config := Configuration{
		Port: "9090",
	}

	oksana := New()
	server := oksana.server(config)

	assert.Equal(t, ":9090", server.Addr)
}
