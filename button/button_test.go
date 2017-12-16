package button

import (
	"io/ioutil"
	"testing"

	"github.com/gernest/chronicles/colors"
	"github.com/gernest/chronicles/styles/theme"

	"github.com/gernest/chronicles/ui/funcs"

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
	ctx.Registry.Register("BaseButton", &Base{})
	u, err := ui.New(`<BaseButton>sign up</BaseButton>`, ctx)
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

func TestStyle(t *testing.T) {
	s := goss.Sheet{
		Class:     make(goss.ClassMap),
		ClassFunc: goss.IDNamer,
	}
	opts := goss.NewOpts()
	opts.FuncMap = funcs.New()
	tm := theme.New(colors.LightContrast)

	err := s.Parse(Style(tm), opts, map[string]interface{}{
		"theme": tm,
	})
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadFile("style.css")
	if err != nil {
		t.Fatal(err)
	}
	e := string(b)
	g := s.Src.String()
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
	// ioutil.WriteFile("style.css", s.Src.Bytes(), 0600)
}
