package ui

import (
	"errors"

	"github.com/broci/chronicles/ui/component"
	"github.com/broci/chronicles/ui/core"
)

type abys struct{}

func (abys) Write(p []byte) (int, error) {
	return 0, nil
}

type UI struct {
	Ctx  *component.Context
	Root *core.Container
}

func New(src interface{}, ctx *component.Context) (*UI, error) {
	u := &UI{Ctx: ctx}
	switch v := src.(type) {
	case *core.Container:
		u.Root = v
		return u, nil
	case component.Component:
		u.Root = &core.Container{
			Kind:      core.Component,
			Component: v,
		}
		return u, nil
	case []byte:
		l, err := core.Parse(v)
		if err != nil {
			return nil, err
		}
		if len(l.Childrens) > 1 {
			return nil, errors.New("There can only be one root element")
		}
		u.Root = l.Childrens[0]
		return u, nil
	case string:
		l, err := core.Parse([]byte(v))
		if err != nil {
			return nil, err
		}
		if len(l.Childrens) > 1 {
			return nil, errors.New("There can only be one root element")
		}
		u.Root = l.Childrens[0]
		return u, nil
	default:
		return nil, errors.New("Unkown source")
	}
}

func (u *UI) Mount() error {
	_, err := u.Root.RenderTo(abys{}, u.Ctx)
	if err != nil {
		return err
	}
	return u.Root.Mount(u.Ctx)
}

func (u *UI) Render() (string, error) {
	_, err := u.Root.RenderTo(abys{}, u.Ctx)
	if err != nil {
		return "", err
	}
	return u.Root.HTML.String(), nil
}
