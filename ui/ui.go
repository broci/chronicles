package ui

import (
	"bytes"
	"errors"
	"html/template"
	"strings"

	"github.com/albrow/vdom"
	"github.com/broci/chronicles/id"
	"github.com/broci/chronicles/ui/component"
	"github.com/broci/chronicles/ui/state"
)

type UI struct {
	Registry *component.Registry
	State    *state.State
	Tree     *vdom.Tree
	root     *coreComponent
}

func New(tpl []byte, r *component.Registry) (*UI, error) {
	u := &UI{
		Registry: r,
		State:    state.New(),
	}
	if err := u.Parse(tpl); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UI) Parse(src []byte) error {
	src = bytes.TrimSpace(src)
	t, err := vdom.Parse(src)
	if err != nil {
		return err
	}
	return u.processTree(t)
}

func (u *UI) processTree(t *vdom.Tree) error {
	if len(t.Children) != 1 {
		return errors.New("There can only be one rootnode")
	}
	root, err := parse(t.Children[0], u.Registry)
	if err != nil {
		return err
	}
	u.root = root
	u.Tree = t
	return nil
}

func (u *UI) Render() error {
	return u.root.Render(u.State, u.Registry)
}

func (u *UI) HTML() (string, error) {
	return u.root.HTML(u.State, u.Registry)
}

var ErrNotRoot = errors.New("root node must be of type *vdom.Element")

func parse(n vdom.Node, r *component.Registry) (*coreComponent, error) {
	e, ok := n.(*vdom.Element)
	if !ok {
		return nil, ErrNotRoot
	}
	attrs := e.AttrMap()
	props := make(map[string]interface{})
	var needs map[string]string
	for k, v := range attrs {
		p, ok := needProp(v)
		if ok {
			needs[k] = p
		} else {
			props[k] = p
		}
	}
	c := &coreComponent{
		node:  n,
		name:  e.Name,
		props: props,
		needs: needs,
		id:    id.Next(),
	}
	ch := e.Children()
	if len(ch) > 0 {
		for _, child := range ch {
			cp, err := parse(child, r)
			if err != nil {
				if err == ErrNotRoot {
					continue
				} else {
					return nil, err
				}
			}
			cp.parent = c
			c.children = append(c.children, cp)
		}
	}
	return c, nil
}

func needProp(p string) (string, bool) {
	p = strings.TrimSpace(p)
	if p == "" {
		return p, false
	}
	if p[0] == '{' && p[len(p)-1] == '}' {
		return p[1 : len(p)-2], true
	}
	return p, false
}

type coreComponent struct {
	name         string
	node         vdom.Node
	parent       *coreComponent
	tplString    string
	tpl          *template.Template
	renderedHTML bytes.Buffer
	props        map[string]interface{}
	id           int64
	needs        map[string]string
	children     []*coreComponent
}

func (c *coreComponent) Template() string {
	return c.tplString
}

func (c *coreComponent) ID() int64 {
	return c.id
}

func (c *coreComponent) Props() map[string]interface{} {
	return c.props
}

func (c *coreComponent) NeedProps() map[string]string {
	return c.needs
}

func (c *coreComponent) Render(s *state.State, r *component.Registry) error {
	tplStr := string(c.node.HTML())
	if c := r.Get(c.name); c != nil {
		tplStr = c.Template()
	}
	for _, child := range c.children {
		h, err := child.HTML(s, r)
		if err != nil {
			return err
		}
		tplStr = strings.Replace(tplStr, string(child.node.HTML()), h, 1)
	}
	ctx := make(map[string]interface{})
	props := make(map[string]interface{})
	if c.parent != nil {
		for k, p := range c.parent.Props() {
			props[k] = p
		}
	}
	for k, v := range c.needs {
		if _, ok := props[v]; !ok {
			npp, ok := s.Get(v)
			if !ok {
				return errors.New("can't find prop " + k)
			}
			props[k] = npp
		}
	}
	for k, v := range c.props {
		props[k] = v
	}
	tpl, err := template.New("component").Parse(tplStr)
	if err != nil {
		return err
	}
	c.tpl = tpl
	ctx["props"] = props
	c.renderedHTML.Reset()
	return c.tpl.Execute(&c.renderedHTML, ctx)
}

func (c *coreComponent) getHTML() {

}

func (c *coreComponent) HTML(s *state.State, r *component.Registry) (string, error) {
	if err := c.Render(s, r); err != nil {
		return "", err
	}
	return c.renderedHTML.String(), nil
}
