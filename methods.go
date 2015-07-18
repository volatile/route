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

// Get makes a route for the GET method.
func Get(pattern string, handler interface{}) {
	Use([]string{"GET"}, pattern, handler)
}

// Post makes a route for the POST method.
func Post(pattern string, handler interface{}) {
	Use([]string{"POST"}, pattern, handler)
}

// Put makes a route for the PUT method.
func Put(pattern string, handler interface{}) {
	Use([]string{"PUT"}, pattern, handler)
}

// Patch makes a route for the PATCH method.
func Patch(pattern string, handler interface{}) {
	Use([]string{"PATCH"}, pattern, handler)
}

// Delete makes a route for the DELETE method.
func Delete(pattern string, handler interface{}) {
	Use([]string{"DELETE"}, pattern, handler)
}

// Head makes a route for the HEAD method.
func Head(pattern string, handler interface{}) {
	Use([]string{"HEAD"}, pattern, handler)
}

// Options makes a route for the OPTIONS method.
func Options(pattern string, handler interface{}) {
	Use([]string{"OPTIONS"}, pattern, handler)
}
