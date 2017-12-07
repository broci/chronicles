package ui

import (
	"bytes"
	"errors"
	"html/template"
	"strings"

	"honnef.co/go/js/dom"

	"github.com/albrow/vdom"
	"github.com/broci/chronicles/id"
	"github.com/broci/chronicles/ui/component"
)

type UI struct {
	Ctx  *component.Context
	Tree *vdom.Tree
	root *coreComponent
	el   dom.Element
}

func New(tpl []byte, ctx *component.Context) (*UI, error) {
	u := &UI{
		Ctx: ctx,
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
	root, err := parse(t.Children[0], u.Ctx)
	if err != nil {
		return err
	}
	u.root = root
	u.Tree = t
	return nil
}

func (u *UI) Render() error {
	return u.root.Render(u.Ctx)
}

func (u *UI) HTML() (string, error) {
	return u.root.HTML(u.Ctx)
}

func (u *UI) Mount(node dom.Element) error {
	return u.root.Mount(u.Ctx)
}

var ErrNotRoot = errors.New("root node must be of type *vdom.Element")

func parse(n vdom.Node, ctx *component.Context) (*coreComponent, error) {
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
			cp, err := parse(child, ctx)
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
	tree         *vdom.Tree
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

func (c *coreComponent) Mount(ctx *component.Context) error {
	h, err := c.HTML(ctx)
	if err != nil {
		return err
	}
	n, err := vdom.Parse([]byte(h))
	if err != nil {
		return err
	}
	patch, err := vdom.Diff(c.tree, n)
	if err != nil {
		return err
	}
	err = patch.Patch(ctx.Element)
	if err != nil {
		return err
	}
	c.tree = n
	return nil
}

func (c *coreComponent) Render(rctx *component.Context) error {
	tplStr := string(c.node.HTML())
	if c := rctx.Registry.Get(c.name); c != nil {
		tplStr = c.Template()
	}
	for _, child := range c.children {
		h, err := child.HTML(rctx)
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
			npp, ok := rctx.State.Get(v)
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
	err = c.tpl.Execute(&c.renderedHTML, ctx)
	if err != nil {
		return err
	}
	newTree, err := vdom.Parse(c.renderedHTML.Bytes())
	if err != nil {
		return err
	}
	c.tree = newTree
	return nil
}

func (c *coreComponent) HTML(ctx *component.Context) (string, error) {
	if err := c.Render(ctx); err != nil {
		return "", err
	}
	return c.renderedHTML.String(), nil
}
