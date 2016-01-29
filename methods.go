package route

import "github.com/volatile/core"

func methodMatch(methods []string, c *core.Context) bool {
	for _, v := range methods {
		if v == c.Request.Method {
			return true
		}
	}
	return false
}

// Get adds a handler to the default handlers stack.
// It executes the handler when request matches GET method and pattern.
func Get(pattern string, handler interface{}) {
	Use([]string{"GET"}, pattern, handler)
}

// Post adds a handler to the default handlers stack.
// It executes the handler when request matches POST method and pattern.
func Post(pattern string, handler interface{}) {
	Use([]string{"POST"}, pattern, handler)
}

// Put adds a handler to the default handlers stack.
// It executes the handler when request matches PUT method and pattern.
func Put(pattern string, handler interface{}) {
	Use([]string{"PUT"}, pattern, handler)
}

// Patch adds a handler to the default handlers stack.
// It executes the handler when request matches PATCH method and pattern.
func Patch(pattern string, handler interface{}) {
	Use([]string{"PATCH"}, pattern, handler)
}

// Delete adds a handler to the default handlers stack.
// It executes the handler when request matches DELETE method and pattern.
func Delete(pattern string, handler interface{}) {
	Use([]string{"DELETE"}, pattern, handler)
}

// Head adds a handler to the default handlers stack.
// It executes the handler when request matches HEAD method and pattern.
func Head(pattern string, handler interface{}) {
	Use([]string{"HEAD"}, pattern, handler)
}

// Options adds a handler to the default handlers stack.
// It executes the handler when request matches OPTIONS method and pattern.
func Options(pattern string, handler interface{}) {
	Use([]string{"OPTIONS"}, pattern, handler)
}
