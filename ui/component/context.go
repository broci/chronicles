package component

import (
	"github.com/broci/chronicles/ui/state"
	"github.com/broci/goss"
	"honnef.co/go/js/dom"
)

type Context struct {
	Mount       bool
	Document    dom.Document
	RootElement dom.Element
	Element     dom.Element
	State       *state.State
	Registry    *Registry
	StyleSheet  *goss.StyleSheet
}

func NewCtx() *Context {
	return &Context{
		State:      state.New(),
		Registry:   NewRegistry(),
		StyleSheet: &goss.StyleSheet{},
	}
}
