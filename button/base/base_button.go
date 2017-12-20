package base

import (
	"honnef.co/go/js/dom"
	"github.com/gernest/chronicles/styles/theme"
	"github.com/gernest/chronicles/ui/component"
	"github.com/gernest/chronicles/ui/state"
	"github.com/gernest/goss"
)

const keyboardFocused = "keyboardFocused"

type Button struct {
	props component.Props
	state *state.State
	button dom.Element
}

func Style(t theme.Theme) goss.CSS {
	return goss.CSS{}
}

func (b *Button) Init(ctx *component.Context) *Button {
	ctx.LocalState.Set(keyboardFocused, false)
	return &Button{
		state: ctx.LocalState,
		props: make(component.Props),
	}
}

func (b *Button) ComponentWillReceiveProps(next component.Props) {
	key:="disabled"
	if  !b.props.Bool(key)
	&& next.Bool(key)&&b.state.Bool(keyboardFocused){
		b.state.Set(keyboardFocused,false)
	}
}


func (b *Button)ComponentDidMount(ctx *component.Context)  {
	b.button=ctx.Element
	util.ListenForFocusKeys(dom.GetWindow())
}