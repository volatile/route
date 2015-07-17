package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/volatile/core"
)

var (
	errUnreachableParams = `route: unreachable named params, use func(c *core.Context, map[string]string) handler type for route "%s %s"`
	errNoParams          = `route: no named params, use a simple func(c *core.Context) handler type for route "%s %s"`
	errUnknownType       = `route: unknown handler type for route "%s %s"`
)

// Use makes a route for the given method.
func Use(method, pattern string, handler interface{}) {
	p := newPat(pattern)

	switch handler.(type) {

	// Raw function
	case func(*core.Context):
		panicUnreachableParams(method, pattern, p)
		use(method, p, func(c *core.Context, _ map[string]string) {
			handler.(func(*core.Context))(c)
		})

	// Raw function with parameters
	case func(*core.Context, map[string]string):
		if !p.hasParams() {
			panic(fmt.Sprintf(errNoParams, method, pattern))
		}
		use(method, p, handler.(func(*core.Context, map[string]string)))

	// String
	case string:
		panicUnreachableParams(method, pattern, p)
		use(method, p, func(c *core.Context, _ map[string]string) {
			c.ResponseWriter.Write([]byte(handler.(string)))
		})

	// []byte
	case []byte:
		panicUnreachableParams(method, pattern, p)
		use(method, p, func(c *core.Context, _ map[string]string) {
			c.ResponseWriter.Write(handler.([]byte))
		})

	// Status
	case int:
		panicUnreachableParams(method, pattern, p)
		use(method, p, func(c *core.Context, _ map[string]string) {
			http.Error(c.ResponseWriter, http.StatusText(handler.(int)), handler.(int))
		})

	// Others
	default:
		switch reflect.ValueOf(handler).Kind() {

		// JSON
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			panicUnreachableParams(method, pattern, p)
			useJSONHandler(method, p, handler)

		// Pointer
		case reflect.Ptr:
			switch reflect.ValueOf(handler).Elem().Kind() {
			// Pointer to JSON
			case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
				panicUnreachableParams(method, pattern, p)
				useJSONHandler(method, p, handler)
			// Pointer to unknown type
			default:
				panic(fmt.Sprintf(errUnknownType, method, pattern))
			}

		// Unknown type
		default:
			panic(fmt.Sprintf(errUnknownType, method, pattern))
		}
	}
}

func useJSONHandler(method string, p *pat, handler interface{}) {
	use(method, p, func(c *core.Context, _ map[string]string) {
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

func use(method string, p *pat, handler func(*core.Context, map[string]string)) {
	core.Use(func(c *core.Context) {
		if c.Request.Method == method && p.match(c) {
			handler(c, p.parseParams(c))
		} else {
			c.Next()
		}
	})
}

func panicUnreachableParams(method, pattern string, p *pat) {
	if p.hasParams() {
		panic(fmt.Sprintf(errUnreachableParams, method, pattern))
	}
}
