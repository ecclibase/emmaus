package oksana

import (
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	id   int
	name string
}

const JSON = `{"id":1,"name":"John Adams"}`

func SampleMethod(context *Context) error {
	return context.JSON(200, "message")
}

func TestContext(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	// Response
	assert.NotNil(t, context.Response)

	// Request
	assert.NotNil(t, context.Request)
}

func TestJSONResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	err := context.JSON(200, User{1, "John Adams"})

	if assert.NoError(t, err) {
		assert.Equal(t, 200, writer.Code)
	}
}

func TestJSONErrorResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	err := context.JSON(200, math.Inf(1))

	assert.Error(t, err)
}

func TestStringResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	err := context.String(200, "this is a test")

	if assert.NoError(t, err) {
		assert.Equal(t, 200, writer.Code)
	}
}

func TestErrorResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	err := context.HTTPError(500, "this is a test")

	if assert.NoError(t, err) {
		assert.Equal(t, 500, writer.Code)
	}
}

func TestAddingParams(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	context.Params = &url.Values{}
	context.Params.Add("key1", "value")
	context.Params.Add("key2", "value")

	assert.Equal(t, context.GetParam("key1"), "value")
	assert.Equal(t, context.GetParam("key2"), "value")
}

func TestSettingHeaders(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	context.SetHeader("Content-Type", "text/html;charset=utf-8")

	assert.Equal(t, context.Response.Header().Get("Content-Type"), "text/html;charset=utf-8")
}

func TestGettingHeaders(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	context.Request.Header.Set("Content-Type", "text/html;charset=utf-8")

	assert.Equal(t, context.GetHeader("Content-Type"), "text/html;charset=utf-8")
}

func TestRedirect(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()

	context := NewContext()
	context.Request = request
	context.Response = NewResponse(writer)

	err := context.Redirect(301, "/")

	if assert.NoError(t, err) {
		assert.Equal(t, 301, writer.Code)
	}
}

func TestHasParam(t *testing.T) {
	oksana := New()

	oksana.GET("/uri", SampleMethod)

	r, _ := http.NewRequest("GET", "/uri?query1=1&query2=2", nil)
	w := httptest.NewRecorder()
	oksana.ServeHTTP(w, r)

	assert.True(t, oksana.Context.HasParam("query1"))
	assert.True(t, oksana.Context.HasParam("query2"))
}

func TestCodeFunc(t *testing.T) {
	oksana := New()

	oksana.GET("/uri", SampleMethod)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	oksana.ServeHTTP(w, r)

	assert.Nil(t, oksana.Context.Code(200))
}
