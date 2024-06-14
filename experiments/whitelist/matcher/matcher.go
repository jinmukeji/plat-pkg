package matcher

import (
	"errors"
	"fmt"

	"github.com/gobwas/glob"
)

type Matcher struct {
	globCache map[string]glob.Glob
}

func NewMatcher() *Matcher {
	return &Matcher{
		globCache: make(map[string]glob.Glob),
	}
}

func (m *Matcher) LoadPatterns(patterns ...string) error {
	for _, p := range patterns {
		g, err := glob.Compile(p)
		if err != nil {
			return err
		}
		m.globCache[p] = g
	}
	return nil
}

func (m *Matcher) Match(pattern, service, method string) (bool, error) {
	g, ok := m.globCache[pattern]
	if !ok {
		return false, errors.New("pattern is not found")
	}
	return matchG(g, service, method)
}

func combineCall(service, method string) string {
	return fmt.Sprintf("%s/%s", service, method)
}

func Match(pattern, service, method string) bool {
	g := glob.MustCompile(pattern)
	match, _ := matchG(g, service, method)
	return match
}

func matchG(g glob.Glob, service, method string) (bool, error) {
	if service == "" {
		return false, errors.New("service should not be empty")
	}
	if method == "" {
		return false, errors.New("method should not be empty")
	}
	s := combineCall(service, method)
	return g.Match(s), nil
}
