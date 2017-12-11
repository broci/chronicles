package core

import (
	"bytes"
	"errors"
	"io"

	"honnef.co/go/js/dom"

	"github.com/albrow/vdom"
	"github.com/broci/chronicles/id"
	"github.com/broci/chronicles/ui/component"
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

	//Elelent is container for html elements i.e anything in tags.
	Element
)

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
	Needs component.Props

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

func ContainerFromNode(e vdom.Node) (*Container, error) {
	c := &Container{
		Node: e,
		ID:   id.Next(),
	}
	switch v := e.(type) {
	case *vdom.Element:
		c.Kind = Element
		props := make(component.Props)
		needs := make(component.Props)
		for k, v := range v.AttrMap() {
			p, ok := component.NeedProp(v)
			if ok {
				needs[k] = p
			} else {
				props[k] = p
			}
		}
		c.Props = props
		c.Needs = needs
		for _, child := range v.Children() {
			ch, err := ContainerFromNode(child)
			if err != nil {
				return nil, err
			}
			c.Children = append(c.Children, ch)
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

func (c *Container) RenderTo(out io.Writer) (int64, error) {
	switch c.Kind {
	case Text, Comment:
		return c.HTML.WriteTo(out)
	default:
		return 0, ErrUnkownNode
	}
}
