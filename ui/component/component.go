package component

import (
	"github.com/broci/chronicles/id"
	"github.com/broci/chronicles/ui/state"
)

type Component interface {
	Template() []byte
	Props() map[string]interface{}
	NeedProps() []string
	SetState(*state.State)
}

type Identity interface {
	ID() string
}

type IDWrap struct {
	id int64
	cm Component
}

func WrapID(c Component) *IDWrap {
	return &IDWrap{id: id.Next(), cm: c}
}

func (i *IDWrap) Template() []byte {
	return i.cm.Template()
}

func (i *IDWrap) Props() map[string]interface{} {
	return i.cm.Props()
}

func (i *IDWrap) NeedProps() []string {
	return i.cm.NeedProps()
}

func (i *IDWrap) SetState(s *state.State) {
	i.cm.SetState(s)
}

func (i *IDWrap) ID() int64 {
	return i.id
}
