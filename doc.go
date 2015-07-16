/*
Package route is a helper for the Core.
It provides syntactic sugar that create a handler filtered by request's method and path.

Usage

Here is the full  example with a classic "Hello, World!", an easy "Hello, World!", a status response, and a JSON encoding:

	package main

	import (
		"fmt"
		"net/http"

		"github.com/volatile/core"
		"github.com/volatile/route"
	)

	func main() {
		route.Get("/classic", func(c *core.Context) {
			fmt.Fprint(c.ResponseWriter, "Hello, World!")
			c.Next()
		})

		route.Get("/", "Hello, World!")

		route.Get("/forbidden", http.StatusForbidden)

		route.Get("/json", &Car {
			ID:    1,
			Brand: "Bentley",
			Model: "Continental GT",
		})

		core.Run()
	}

	type Car struct {
		ID    int    `json:"id"`
		Brand string `json:"brand"`
		Model string `json:"model"`
	}

Method filtering

Functions exists for the most common and standard HTTP methods.
If you need to handle a custom method, use the router.Use function.

Handlers

Handlers can be of different types for the best readability and without losing performance…

Raw body

You can use a string or a []byte to send a raw text or the result of a rendering function that returns a raw body ready to be sent.

Status code

You can provide an int to just send a status code.

JSON

You can provide a struct, a map, a slice or an array that will be marshalled and sent as JSON.

Classic function

Obviously, a classic func(c *core.Context) can be used for more flexibility or if you need to use c.Next() inside the handler.
*/
package route
