package main

import (
	"github.com/broci/chronicles/ui"
	"github.com/broci/chronicles/ui/component"
	"honnef.co/go/js/dom"
)

func main() {
	doc := dom.GetWindow().Document()
	ctx := component.NewCtx()
	ctx.Document = doc
	u, err := ui.New(`<h1> hello, world </h1>`, ctx)
	if err != nil {
		panic(err)
	}
	if err = u.Mount(); err != nil {
		panic(err)
	}
}
