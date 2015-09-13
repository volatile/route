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

// Use makes a route for the given methods.
func Use(methods []string, pattern string, handler interface{}) {
	p := newPat(pattern)

	switch handler.(type) {

	// Context handler
	case func(*core.Context):
		if len(p.namedParams) > 0 {
			panic(fmt.Sprintf(errUnreachableParams, methods, pattern))
		}
		use(methods, p, func(c *core.Context, _ map[string]string) {
			handler.(func(*core.Context))(c)
		})

	// Context handler with parameters
	case func(*core.Context, map[string]string):
		if len(p.namedParams) == 0 {
			panic(fmt.Sprintf(errNoParams, methods, pattern))
		}
		use(methods, p, handler.(func(*core.Context, map[string]string)))

	// Unknown type
	default:
		panic(fmt.Sprintf(errUnknownType, methods, pattern))

	}
}

func use(methods []string, p *pat, handler func(*core.Context, map[string]string)) {
	core.Use(func(c *core.Context) {
		if methodMatch(methods, c) && p.match(c) {
			handler(c, p.parseParams(c))
		} else {
			c.Next()
		}
	})
}
