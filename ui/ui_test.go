package ui

import (
	"testing"

	"github.com/broci/chronicles/ui/component"
)

type dummy struct {
	tpl string
}

func (d dummy) Template() string {
	return d.tpl
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
				{"todo", dummy{tpl: `<h2>hello, world</h2>`}},
			},
			`<h2>hello, world</h2>`,
			"Component only",
		},
		{
			`<todo name="gernest"></todo>`,
			[]struct {
				name string
				cmp  component.Component
			}{
				{"todo", dummy{tpl: `<h2>hello, world {{.props.name}}</h2>`}},
			},
			`<h2>hello, world gernest</h2>`,
			"Component with props",
		},
	}

	for _, s := range sample {
		t.Run(s.desc, func(ts *testing.T) {
			r := component.NewRegistry()
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
