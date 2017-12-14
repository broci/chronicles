package core

import (
	"reflect"
	"testing"

	"github.com/gernest/chronicles/ui/component"
)

type tCase struct {
	name string
	kind Kind
}

func TestContainer_Parse(t *testing.T) {
	sample := []struct {
		src      string
		root     string
		rootType Kind
		props    component.Props
		children []tCase
	}{
		{`<h1>hello,world</h1>`, `h1`, Element, component.Props{},
			[]tCase{
				{"", Text},
			}},
		{`<demo>hello,world</demo>`, `demo`, Element, component.Props{},
			[]tCase{
				{"", Text},
			}},
		{`<demo key="value">hello,world</demo>`, `demo`, Element, component.Props{
			"key": "value",
		},
			[]tCase{
				{"", Text},
			}},
	}

	for _, v := range sample {
		l, err := Parse([]byte(v.src))
		if err != nil {
			t.Fatal(err)
		}
		if len(l.Childrens) > 1 {
			t.Errorf("expected one root element got %d", len(l.Childrens))
		}
		node := l.Childrens[0]
		if node.Name != v.root {
			t.Errorf("expected %s got %s", v.root, node.Name)
		}
		if node.Kind != v.rootType {
			t.Errorf("expected %v got %v", v.rootType, node.Kind)
		}
		for k, v := range v.props {
			nv := node.Props[k]
			if !reflect.DeepEqual(v, nv) {
				t.Errorf("expected %v got %v", v, nv)
			}
		}
		for i, ch := range node.Children {
			vch := v.children[i]
			if ch.Name != vch.name {
				t.Errorf("expected %s got %s", vch.name, ch.Name)
			}
			if ch.Kind != vch.kind {
				t.Errorf("expected %v got %v", ch.Kind, vch.kind)
			}
		}
	}
}
