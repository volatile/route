package route

import (
	"regexp"

	"github.com/volatile/core"
)

// pat represents a pattern with a regexp for filtering and the list of named params and their position.
type pat struct {
	regexp      *regexp.Regexp
	namedParams map[string]int
}

// newPat compiles the pattern regex and saves params with their name and position.
func newPat(pattern string) *pat {
	p := &pat{
		regexp: regexp.MustCompile(pattern),
	}

	namedParams := p.regexp.SubexpNames()[1:]
	if len(namedParams) > 0 {
		p.namedParams = make(map[string]int)
		for i, v := range namedParams {
			if v != "" {
				p.namedParams[v] = i
			}
		}
	}

	return p
}

func (p *pat) match(c *core.Context) bool {
	return p.regexp.MatchString(c.Request.URL.String())
}

func (p *pat) parseParams(c *core.Context) map[string]string {
	if len(p.namedParams) > 0 {
		params := make(map[string]string)
		paramsValues := p.regexp.FindStringSubmatch(c.Request.URL.String())[1:]
		for k, v := range p.namedParams {
			params[k] = paramsValues[v]
		}
		return params
	}
	return nil
}
