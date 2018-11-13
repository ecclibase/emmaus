package oksana

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCreation(t *testing.T) {
	service := New()
	group := service.Group("/group")

	assert.True(t, reflect.TypeOf(group).String() == "*oksana.Group")
}

func TestGroupGetRoute(t *testing.T) {
	service := New()
	group := service.Group("/group")

	group.GET("/path", func(context *Context) error {
		return context.String(200, "This is a test")
	})

	route := service.Router.find("GET", "/path")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
}

func TestGroupPostRoute(t *testing.T) {
	service := New()
	group := service.Group("/group")

	group.POST("/path", func(context *Context) error {
		return context.String(200, "This is a test")
	})

	route := service.Router.find("POST", "/path")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
}

func TestGroupPatchRoute(t *testing.T) {
	service := New()
	group := service.Group("/group")

	group.PATCH("/path", func(context *Context) error {
		return context.String(200, "This is a test")
	})

	route := service.Router.find("PATCH", "/path")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
}

func TestGroupPutRoute(t *testing.T) {
	service := New()
	group := service.Group("/group")

	group.PUT("/path", func(context *Context) error {
		return context.String(200, "This is a test")
	})

	route := service.Router.find("PUT", "/path")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
}

func TestGroupDeleteRoute(t *testing.T) {
	service := New()
	group := service.Group("/group")

	group.DELETE("/path", func(context *Context) error {
		return context.String(200, "This is a test")
	})

	route := service.Router.find("DELETE", "/path")

	assert.True(t, reflect.TypeOf(route).String() == "oksana.Route")
}
