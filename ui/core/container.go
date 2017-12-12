package core

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"strings"

	"honnef.co/go/js/dom"

	"github.com/albrow/vdom"
	"github.com/broci/chronicles/id"
	"github.com/broci/chronicles/ui/component"
	"github.com/broci/chronicles/ui/funcs"
	"github.com/broci/goss"
)

// ErrUnkownNode is raised for unsupported node types.
var ErrUnkownNode = errors.New("ui: Unknown node")

// Kind defines a type of Container.
type Kind uint32

const (

	// Comment is a container for html comments
	Comment Kind = iota + 1

	// Text is a container for html text nodes i.e anything not in tags.
	Text

	//Element is container for html elements i.e anything in tags.
	Element

	// Component is any non standard html tag.
	Component
)

func (k Kind) String() string {
	switch k {
	case Comment:
		return "COMMENT"
	case Text:
		return "TEXT"
	case Element:
		return "ELEMENT"
	default:
		return ""
	}
}

// Container represents the parsed html node.
type Container struct {

	// Kind represent the type of the container
	Kind Kind

	// Is tha name of the container element as per vdom.Node.Name
	Name string

	// Node is the original node as provided by the template. This means pre
	// rendered html node from the components templates.
	Node vdom.Node

	// Parent is the pointer to the parent of this container. If it is nil then
	// this is the root container.
	Parent *Container

	//This is the dom node after rendering is complete.
	RenderedNode vdom.Node

	// HTML is the buffer of rendered html. It contains the same content as
	// RenderedNode.HTML().
	HTML bytes.Buffer

	Props component.Props

	//Unique ID of this container. This is for internal use only.
	ID int64

	// Needs properties need by this component to render. If they are missing this
	// component will never be rendered.
	//
	// They will be searched from parent first, if missing then they will be
	// searched from the global state.
	Needs []string

	// Sheet is the container style. It will be passed to the template when
	// rendering.
	Sheet *goss.Sheet

	// Element is the actual dom element that is atached to this container.
	Element dom.Element

	Children []*Container
}

// List is a group of containers.
type List struct {
	Tree      *vdom.Tree
	Childrens []*Container
}

func Parse(src []byte) (*List, error) {
	t, err := vdom.Parse(src)
	if err != nil {
		return nil, err
	}
	l := &List{Tree: t}
	for _, e := range t.Children {
		c, err := ContainerFromNode(e)
		if err != nil {
			return nil, err
		}
		l.Childrens = append(l.Childrens, c)
	}
	return l, nil
}

// ContainerFromNode creates a container out of the node e.
func ContainerFromNode(e vdom.Node) (*Container, error) {
	c := &Container{
		Node: e,
		ID:   id.Next(),
	}
	switch v := e.(type) {
	case *vdom.Element:
		c.Kind = Element
		props := make(component.Props)
		needs := []string{}
		for k, v := range v.AttrMap() {
			p, ok := component.NeedProp(v)
			if ok {
				needs = append(needs, p)
			} else {
				props[k] = p
			}
		}
		c.Props = props
		c.Needs = needs
		c.Name = v.Name
		for _, child := range v.Children() {
			ch, err := ContainerFromNode(child)
			if err != nil {
				return nil, err
			}
			c.Children = append(c.Children, ch)
		}

		// anything other than standard html tags is a component.
		if !goss.IsHTMLTAG(c.Name) {
			c.Kind = Component
		}
	case *vdom.Text:
		c.Kind = Text
		c.HTML.Write(v.Value)
	case *vdom.Comment:
		c.Kind = Comment
		c.HTML.Write(v.Value)
	default:
		return nil, ErrUnkownNode
	}
	return c, nil
}

func (c *Container) RenderTo(out io.Writer, ctx *component.Context) (int64, error) {
	switch c.Kind {
	case Text, Comment:
		return c.HTML.WriteTo(out)
	case Element, Component:
		tplStr := string(c.Node.HTML())
		var needs []string
		props := make(component.Props)
		if c.Parent != nil {
			for k, p := range c.Parent.Props {
				props[k] = p
			}
		}
		if c.Kind == Component {
			if cmp := ctx.Registry.Get(c.Name); cmp != nil {
				tplStr = cmp.Template()

				if cp, ok := cmp.(component.HasProps); ok {
					for k, v := range cp.Props() {
						props[k] = v
					}
				}
				if cp, ok := cmp.(component.NeedsProps); ok {
					for _, v := range cp.NeedsProps() {
						needs = append(needs, v)
					}
				}
			}
		}

		var buf bytes.Buffer
		for _, child := range c.Children {
			buf.Reset()
			_, err := child.RenderTo(&buf, ctx)
			if err != nil {
				return 0, err
			}
			tplStr = strings.Replace(tplStr, string(child.Node.HTML()), buf.String(), 1)
		}
		if c.Parent != nil {
			for k, p := range c.Parent.Props {
				props[k] = p
			}
		}
		needs = append(needs, c.Needs...)
		for _, v := range needs {
			if _, ok := props[v]; !ok {
				npp, ok := ctx.State.Get(v)
				if !ok {
					return 0, errors.New("can't find prop " + v)
				}
				props[v] = npp
			}
		}
		props["parent"] = c.Props
		props["classes"] = c.Sheet.Class
		tpl, err := template.New("component").Funcs(funcs.New()).Parse(tplStr)
		if err != nil {
			return 0, err
		}
		c.HTML.Reset()
		err = tpl.Execute(&c.HTML, props)
		if err != nil {
			return 0, err
		}
		t, err := vdom.Parse(c.HTML.Bytes())
		if err != nil {
			return 0, err
		}
		c.RenderedNode = t.Children[0]
		return int64(c.HTML.Len()), nil
	default:
		return 0, ErrUnkownNode
	}
}
