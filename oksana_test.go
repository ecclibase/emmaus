package oksana

import (
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
	service := New()
	assert.NotNil(t, service)
}
