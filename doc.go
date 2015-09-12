/*
Package route is a helper for the Core (https://github.com/volatile/core).
It provides syntactic sugar that create a handler filtered by request's method and path.

Installation

In the terminal:

	$ go get github.com/volatile/route

Usage

Example:

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

Method filtering

Functions exists for the most common and standard HTTP methods.
If you need to handle custom methods, use the Use function with the methods in a strings slice as the first parameter.

Remember that HTTP methods are case-sensitive. See RFC 7231 4.1 (https://tools.ietf.org/html/rfc7231#section-4.1).

Path filtering

A regular expression is used to match the request path.
So you keep a full control over your routing strategies.
We think the regular expressions offer the best balance between performance and power for this kind of job.

If you need to use named parameters from the URL, just use a regexp named group like (?P<id>[0-9]+) and a func(c *core.Context, map[string]string) handler type.
*/
package route
