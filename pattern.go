package route

import (
	"regexp"

	"github.com/volatile/core"
)

// pat represents pattern with a regexp for filtering and the list of named params and their position.
type pat struct {
	regexp      *regexp.Regexp
	paramsNames map[string]int
}

func newPat(pattern string) *pat {
	p := &pat{
		regexp: regexp.MustCompile(pattern),
	}

	paramsNames := p.regexp.SubexpNames()[1:]
	if len(paramsNames) > 0 {
		p.paramsNames = make(map[string]int)
		for i, v := range paramsNames {
			if v != "" {
				p.paramsNames[v] = i
			}
		}
	}

	return p
}

func (p *pat) match(c *core.Context) bool {
	return p.regexp.MatchString(c.Request.URL.String())
}

func (p *pat) hasParams() bool {
	return p.paramsNames != nil
}

func (p *pat) parseParams(c *core.Context) map[string]string {
	if p.paramsNames != nil {
		params := make(map[string]string)
		paramsValues := p.regexp.FindStringSubmatch(c.Request.URL.String())[1:]
		for k, v := range p.paramsNames {
			params[k] = paramsValues[v]
		}
		return params
	}
	return nil
}
