package text

import (
	"github.com/gernest/chronicles/ui/component"
)

type Text struct {
	Text string
}

func (t *Text) Template() string {
	return t.Text
}

func (t *Text) Init(*component.Context) component.Component {
	return t
}
