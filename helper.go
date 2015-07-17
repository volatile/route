package route

import (
	"fmt"

	"github.com/volatile/core"
)

var (
	errUnreachableParams = `route: unreachable named parameters, use func(c *core.Context, map[string]string) handler type for route "%s %s"`
	errNoParams          = `route: no named parameters, use a simple func(c *core.Context) handler type for route "%s %s"`
	errUnknownType       = `route: unknown handler type for route "%s %s"`
)

// Use makes a route for the given method.
func Use(method, pattern string, handler interface{}) {
	p := newPat(pattern)

	switch handler.(type) {

	// Context handler
	case func(*core.Context):
		if p.namedParams != nil {
			panic(fmt.Sprintf(errUnreachableParams, method, pattern))
		}
		use(method, p, func(c *core.Context, _ map[string]string) {
			handler.(func(*core.Context))(c)
		})

	// Context handler with parameters
	case func(*core.Context, map[string]string):
		if p.namedParams == nil {
			panic(fmt.Sprintf(errNoParams, method, pattern))
		}
		use(method, p, handler.(func(*core.Context, map[string]string)))

	// Unknown type
	default:
		panic(fmt.Sprintf(errUnknownType, method, pattern))

	}
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
