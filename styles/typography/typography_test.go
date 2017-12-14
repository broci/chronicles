package typography

import (
	"testing"

	"github.com/gernest/chronicles/styles/unit"

	"github.com/gernest/chronicles/styles/palette"
)

func TestTypography(t *testing.T) {
	t.Run("With default material design spec", func(ts *testing.T) {
		g := New(palette.Palette{})
		s := unit.Format(g.FontSize)
		e := "14px"
		if s != e {
			ts.Errorf("expected %s got %s", e, s)
		}
	})

	t.Run("With custom fontsize", func(ts *testing.T) {
		g := New(palette.Palette{}, Opts{FontSize: 15})
		s := unit.Format(g.FontSize)
		e := "15px"
		if s != e {
			ts.Errorf("expected %s got %s", e, s)
		}
	})
	t.Run("Display fontsize", func(ts *testing.T) {
		g := New(palette.Palette{}, Opts{HTMLFontSize: 10})
		s := unit.Format(g.Style[Display4].FontSize)
		e := "11rem"
		if s != e {
			ts.Errorf("expected %s got %s", e, s)
		}
	})
}
