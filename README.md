[![Build Status](https://travis-ci.com/ecclibase/oksana.svg?branch=master)](https://travis-ci.com/ecclibase/oksana)
[![codecov](https://codecov.io/gh/ecclibase/oksana/branch/master/graph/badge.svg)](https://codecov.io/gh/ecclibase/oksana)
[![GoDoc](https://godoc.org/github.com/olebedev/config?status.png)](https://godoc.org/github.com/ecclibase/oksana)
[![Go Report Card](https://goreportcard.com/badge/github.com/ecclibase/oksana)](https://goreportcard.com/report/github.com/ecclibase/oksana)

# Oksana

Microservice framework written in Go.

## Example

```go
func main() {
    o := oksana.New()

    // index handler
    o.GET("/", func(c oksana.Context) error {
        return c.JSON(200, "Hello World!")
    })

    o.Start()
}
```

## Config

Oksana provides the ability to configure your service from a central config file.
Copy `.env.example` to `.env` into the source of your service. You can place any
environment variables into this file. The `.env` is loaded into a Config struct
which can be read from anywhere in the service.

## Routes

### Add Routes

```go
// Add GET route
o.GET('/endpoint', handler)

// Add DELETE route
o.DELETE('/endpoint', handler)

// Add PATCH route
o.PATCH('/endpoint', handler)

// Add POST route
o.POST('/endpoint', handler)

// Add PUT route
o.PUT('/endpoint', handler)
```

## Route Groups

```go
// Create Route Group
g := o.Group('/prefix')
g.GET('/endpoint', handler)  // GET /prefix/endpoint
g.POST('/endpoint', handler) // POST /prefix/endpoint
```

### Middleware

Groups can have defined middleware. These middleware handlers will be executed for every route within the group:

```go
g := o.Group('/prefix', middlewareHandler1, middlewareHandler1)
```

Groups can have middleware handlers that are executed for every route within the group:

```go
g := o.Group('/prefix')
g.Middleware(middlwareHandler1)
g.Middleware(middlwareHandler2)
```

If you need middleware to be executed on specific routes you can add middleware to the route definition:

```go
g := o.Group('/prefix')
g.GET('/endpoint', handler, middleware)
```

## Development

Oksana uses golang's [dep](https://github.com/golang/dep) for dependency management. Make sure dep is installed on your local development machine.

To pull in the required dependencies, run the following command: `dep ensure`.