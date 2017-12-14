package component

import (
	"github.com/gernest/chronicles/ui/state"
	"github.com/gernest/goss"
	"honnef.co/go/js/dom"
)

type Context struct {
	Mount       bool
	Document    dom.Document
	RootElement dom.Element
	Element     dom.Element

	// State is the global state.
	State *state.State

	// LocalState is the state that only works inside the component boundary.This
	// is not passed to child components, as state is private to each component.
	//
	//If you want values to be accesible to anyone use State field.
	LocalState *state.State
	Registry   *Registry
	StyleSheet *goss.StyleSheet
}

func NewCtx() *Context {
	return &Context{
		State:      state.New(),
		Registry:   NewRegistry(),
		StyleSheet: &goss.StyleSheet{},
	}
}
