package ui

import (
	"testing"

	"github.com/broci/chronicles/ui/component"
	"github.com/broci/chronicles/ui/registry"
)

type todo struct{}

func (todo) Template() string {
	return `<h2>hello, world</h2>`
}

func TestUI_Parse(t *testing.T) {
	sample := []struct {
		src        string
		components []struct {
			name string
			cmp  component.Component
		}
		expect string
		desc   string
	}{
		{
			`<todo></todo>`,
			[]struct {
				name string
				cmp  component.Component
			}{
				{"todo", todo{}},
			},
			`<h2>hello, world</h2>`,
			"Component only",
		},
	}

	for _, s := range sample {
		t.Run(s.desc, func(ts *testing.T) {
			r := registry.New()
			for _, c := range s.components {
				r.Register(c.name, c.cmp)
			}
			u, err := New([]byte(s.src), r)
			if err != nil {
				ts.Fatal(err)
			}
			g, err := u.HTML()
			if err != nil {
				ts.Fatal(err)
			}
			if g != s.expect {
				t.Errorf("expected %s got %s", s.expect, g)
			}
		})
	}
}
