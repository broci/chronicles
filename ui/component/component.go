package component

import (
	"strings"

	"github.com/gernest/chronicles/styles/theme"
	"github.com/gernest/goss"
	"honnef.co/go/js/dom"
)

type Component interface {
	Template() string
	Init(*Context) Component
}

type HasProps interface {
	Props() Props
}

type NeedsProps interface {
	NeedsProps() []string
}

type CanListen interface {
	ListenTo() EventListers
}

type DidMount interface {
	ComponentDidMount(*Context)
}

type WillMount interface {
	ComponentWillMount(*Context)
}

type HasStyle interface {
	ComponentStyle(theme.Theme) goss.CSS
}

type WillRecieveProps interface {
	ComponentWillReceiveProps(props Props)
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

type EventListers map[string]func(dom.Event)

func MergeProps(props ...Props) Props {
	p := make(Props)
	for _, pp := range props {
		for k, v := range pp {
			p[k] = v
		}
	}
	return p
}

// Bool finds prop with key, and casts the result to bool. This will return
// false when the key is not set.
func (p Props) Bool(key string) bool {
	v, ok := p[key]
	if !ok {
		return false
	}
	return v.(bool)
}
