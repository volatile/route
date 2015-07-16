package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/volatile/core"
)

var errUnknownType = `route: unknown handler type for route "%s %s"`

func setHandler(method, path string, handler interface{}) {
	switch handler.(type) {
	// Raw function
	case func(*core.Context):
		Use(method, path, handler.(func(*core.Context)))
	// Status
	case int:
		Use(method, path, func(c *core.Context) {
			http.Error(c.ResponseWriter, http.StatusText(handler.(int)), handler.(int))
		})
	// String
	case string:
		Use(method, path, func(c *core.Context) {
			c.ResponseWriter.Write([]byte(handler.(string)))
		})
	// []byte
	case []byte:
		Use(method, path, func(c *core.Context) {
			c.ResponseWriter.Write(handler.([]byte))
		})
	// Others
	default:
		switch reflect.ValueOf(handler).Kind() {
		// JSON
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			useJSONHandler(method, path, handler)
		// Pointer
		case reflect.Ptr:
			switch reflect.ValueOf(handler).Elem().Kind() {
			// Pointer to JSON
			case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
				useJSONHandler(method, path, handler)
			// Pointer to unknown type
			default:
				panic(fmt.Sprintf(errUnknownType, method, path))
			}
		// Unknown type
		default:
			panic(fmt.Sprintf(errUnknownType, method, path))
		}
	}
}

func useJSONHandler(method, path string, handler interface{}) {
	Use(method, path, func(c *core.Context) {
		c.ResponseWriter.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(handler)
		if err != nil {
			log.Println(err)
			http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		c.ResponseWriter.Write(js)
	})
}

func isValidStatus(v int) bool {
	return v >= 200 && v <= 599
}

// Get makes a route for the GET method.
func Get(path string, handler interface{}) {
	setHandler("GET", path, handler)
}

// Post makes a route for the POST method.
func Post(path string, handler interface{}) {
	setHandler("POST", path, handler)
}

// Put makes a route for the PUT method.
func Put(path string, handler interface{}) {
	setHandler("PUT", path, handler)
}

// Patch makes a route for the PATCH method.
func Patch(path string, handler interface{}) {
	setHandler("PATCH", path, handler)
}

// Delete makes a route for the DELETE method.
func Delete(path string, handler interface{}) {
	setHandler("DELETE", path, handler)
}

// Head makes a route for the HEAD method.
func Head(path string, handler interface{}) {
	setHandler("HEAD", path, handler)
}

// Options makes a route for the OPTIONS method.
func Options(path string, handler interface{}) {
	setHandler("OPTIONS", path, handler)
}

// Use makes a route for the given method.
func Use(method, path string, handler func(*core.Context)) {
	core.Use(func(c *core.Context) {
		if c.Request.Method == method && c.Request.URL.String() == path {
			handler(c)
		} else {
			c.Next()
		}
	})
}
