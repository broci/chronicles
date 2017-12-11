package component

import "strings"

type Component interface {
	Template() string
}

type Identity interface {
	ID() string
}

type Props map[string]interface{}

func NeedProp(p string) (string, bool) {
	p = strings.TrimSpace(p)
	if p == "" {
		return p, false
	}
	if p[0] == '{' && p[len(p)-1] == '}' {
		return p[1 : len(p)-2], true
	}
	return p, false
}
