package button

import (
	"strings"
	"testing"

	"github.com/gernest/chronicles/text"
	"github.com/gernest/chronicles/ui"
	"github.com/gernest/chronicles/ui/component"
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
func TestBase_Props(t *testing.T) {
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
	if !strings.Contains(v, "root") {
		t.Error("expected class to be set")
	}
}
