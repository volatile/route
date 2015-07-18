<p align="center"><img src="https://cloud.githubusercontent.com/assets/9503891/8712746/59aa7b40-2b60-11e5-9d79-fbfdafd21d9c.png" alt="Volatile Route" title="Volatile Route"><br><br></p>

Volatile Route is a helper for the [Core](https://github.com/volatile/core).  
It provides syntactic sugar that create a handler filtered by request's method and path.

## Installation

```Shell
$ go get -u github.com/volatile/route
```

## Usage

```Go
package main

import (
	"fmt"

	"github.com/volatile/core"
	"github.com/volatile/route"
)

func main() {
	route.Get("^/$", func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

	// Named parameters
	route.Get("^/(?P<name>[A-Za-z]+)$", func(c *core.Context, params map[string]string) {
		fmt.Fprintf(c.ResponseWriter, "Hello, %s!", params["name"])
	})

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/route?status.svg)](https://godoc.org/github.com/volatile/route)

### Method filtering

Functions exists for the most common and standard HTTP methods.  
If you need to handle a custom methods, use the `router.Use` function with the methods in a strings slice as the first parameter.

### Path filtering

A [regular expression](https://golang.org/pkg/regexp/syntax/) is used to match the request path.  
So you keep a full control over your routing strategies.  
We think the regular expressions offer the best balance between performance and power for this kind of job.

If you need to use named parameters from the URL, just use a regexp named group like `(?P<id>[0-9]+)` and a `func(c *core.Context, map[string]string)` handler type.
