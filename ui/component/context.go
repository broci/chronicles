package component

import (
	"github.com/broci/chronicles/ui/state"
	"honnef.co/go/js/dom"
)

type Context struct {
	Mount       bool
	Document    dom.Document
	RootElement dom.Element
	Element     dom.Element
	State       *state.State
	Registry    *Registry
}

func NewCtx() *Context {
	return &Context{
		State:    state.New(),
		Registry: NewRegistry(),
	}
}
