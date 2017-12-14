package button

import (
	"testing"

	"github.com/gernest/chronicles/text"
	"github.com/gernest/chronicles/ui"
	"github.com/gernest/chronicles/ui/component"
	"github.com/gernest/goss"
)

func TestBase(t *testing.T) {
	ctx := component.NewCtx()
	ctx.StyleSheet.Namer = goss.IDNamer
	ctx.Registry.Register("BaseButton", &Base{})
	u, err := ui.New(`<BaseButton/>`, ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := u.Render()
	if err != nil {
		t.Fatal(err)
	}
	e := `<button class="root-id"></button>`
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}
func TestBase_Children(t *testing.T) {
	ctx := component.NewCtx()
	ctx.StyleSheet.Namer = goss.IDNamer
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
	e := `<button class="root-id">sign up</button>`
	if v != e {
		t.Errorf("expected %s got %s", e, v)
	}
}
