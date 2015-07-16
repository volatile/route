package route

// Get makes a route for the GET method.
func Get(pattern string, handler interface{}) {
	Use("GET", pattern, handler)
}

// Post makes a route for the POST method.
func Post(pattern string, handler interface{}) {
	Use("POST", pattern, handler)
}

// Put makes a route for the PUT method.
func Put(pattern string, handler interface{}) {
	Use("PUT", pattern, handler)
}

// Patch makes a route for the PATCH method.
func Patch(pattern string, handler interface{}) {
	Use("PATCH", pattern, handler)
}

// Delete makes a route for the DELETE method.
func Delete(pattern string, handler interface{}) {
	Use("DELETE", pattern, handler)
}

// Head makes a route for the HEAD method.
func Head(pattern string, handler interface{}) {
	Use("HEAD", pattern, handler)
}

// Options makes a route for the OPTIONS method.
func Options(pattern string, handler interface{}) {
	Use("OPTIONS", pattern, handler)
}
