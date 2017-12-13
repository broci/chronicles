package button

import (
	"testing"

	"github.com/broci/chronicles/text"
	"github.com/broci/chronicles/ui"
	"github.com/broci/chronicles/ui/component"
)

func TestBase(t *testing.T) {
	ctx := component.NewCtx()
	ctx.Registry.Register("BaseButton", &Base{})
	u, err := ui.New(`<BaseButton/>`, ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := u.Render()
	if err != nil {
		t.Fatal(err)
	}
	e := "<button></button>"
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}
func TestBase_Children(t *testing.T) {
	ctx := component.NewCtx()
	ctx.Registry.Register("BaseButton", &Base{
		Children: &text.Text{Text: "sign up"},
	})
	u, err := ui.New(`<BaseButton/>`, ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := u.Render()
	if err != nil {
		t.Fatal(err)
	}
	e := "<button>sign up</button>"
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}
