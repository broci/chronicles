package ui

import (
	"errors"

	"github.com/gernest/chronicles/ui/component"
	"github.com/gernest/chronicles/ui/core"
)

type abys struct{}

func (abys) Write(p []byte) (int, error) {
	return 0, nil
}

// UI central user interface rendering implementation. This follows a mix of
// ideas from many fronend libraries,
// mostly react and vue.
type UI struct {
	Ctx               *component.Context
	Root              *core.Container
	changedContainers chan *core.Container
	initialized       bool
}

// New initializes and returns a new *UI object from src.
//
//Src can be one of,
//	*core.Container
//	component.Component
//	[]bute
//	string
func New(src interface{}, ctx *component.Context) (*UI, error) {
	u := &UI{
		Ctx:               ctx,
		changedContainers: make(chan *core.Container, 1000),
	}
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

// Mount this starts UI rendering logic and mounts the components.This only need
// to be called once, any subsequent calls have no effect.
func (u *UI) Mount() error {
	u.Init()
	_, err := u.Root.RenderTo(abys{}, u.Ctx)
	if err != nil {
		return err
	}
	return u.Root.Mount(u.Ctx)
}

// Render generates html without mounting the components to dom.
func (u *UI) Render() (string, error) {
	_, err := u.Root.RenderTo(abys{}, u.Ctx)
	if err != nil {
		return "", err
	}
	return u.Root.HTML.String(), nil
}

// Init initializes goroutine for rerendering.
func (u *UI) Init() {
	if u.initialized {
		return
	}
	go func() {
		for {
			select {
			case c := <-u.changedContainers:
				go c.Rerender(u.Ctx)
			}
		}
	}()
	u.initialized = true
}
