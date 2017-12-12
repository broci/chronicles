package component

import "strings"

type Component interface {
	Template() string
}

type HasProps interface {
	Props() Props
}

type NeedsProps interface {
	NeedsProps() []string
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
	i := strings.Index(p, "{{")
	if i != -1 {
		e := strings.Index(p, "}}")
		if e != -1 {
			txt := p[i:e]
			txt = strings.TrimSpace(txt)
			if !strings.HasPrefix(txt, ".") {
				return "", false
			}
			return txt[1:], true
		}
	}
	return p, false
}
