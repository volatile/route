/*
Package route is a helper for the core (https://godoc.org/github.com/volatile/core).
It provides syntactic sugar that wraps a handler with request method and path filtering.

Method filtering

Functions exists for the most common and standard HTTP methods.

Example for a GET route:

	route.Get("^/$", func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

If you need to handle custom methods, Use receives the methods in a strings slice:

	route.Use([]string{"GET", "POST"}, "^/$", func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World from GET or POST!")
	})

Remember that HTTP methods are case-sensitive.
See RFC 7231 4.1 (https://tools.ietf.org/html/rfc7231#section-4.1).

Path filtering

A regular expression is used to match the request path.
We think it offers the best balance between performance and power for this kind of job.

If you need to use named parameters from the URL, just use a regexp named group and a func(*core.Context, map[string]string) handler type:

	route.Get("^/(?P<name>[A-Za-z]+)$", func(c *core.Context, params map[string]string) {
		fmt.Fprintf(c.ResponseWriter, "Hello, %s!", params["name"])
	})
*/
package route
