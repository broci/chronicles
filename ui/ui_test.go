package ui

import (
	"testing"

	"github.com/gernest/chronicles/ui/component"
)

type dummy struct {
	tpl string
}

func (d dummy) Init() component.Component {
	return d
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
				{"todo", dummy{tpl: `<h2>hello, world {{.name}}</h2>`}},
			},
			`<h2>hello, world gernest</h2>`,
			"Component with props",
		},
	}

	for _, s := range sample {
		t.Run(s.desc, func(ts *testing.T) {
			ctx := component.NewCtx()
			for _, c := range s.components {
				ctx.Registry.Register(c.name, c.cmp)
			}

			u, err := New([]byte(s.src), ctx)
			if err != nil {
				ts.Fatal(err)
			}
			g, err := u.Render()
			if err != nil {
				ts.Fatal(err)
			}
			if g != s.expect {
				t.Errorf("expected %s got %s", s.expect, g)
			}
		})
	}
}
