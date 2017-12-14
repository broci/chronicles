package mixins

import (
	"github.com/gernest/chronicles/styles/breakpoint"
	"github.com/gernest/chronicles/styles/spacing"
	"github.com/gernest/chronicles/styles/unit"
	"github.com/gernest/goss"
)

type Mixins struct {
	ToolBar goss.CSS
	bp      breakpoint.Breakpoints
	sp      spacing.Spacing
}

func New(bp breakpoint.Breakpoints, sp spacing.Spacing) Mixins {
	m := Mixins{bp: bp, sp: sp}
	m.ToolBar = goss.CSS{
		goss.MinHeight: unit.Format(56),
		m.bp.Up(breakpoint.XS) + "and (orientation: landscape)": goss.CSS{
			goss.MinHeight: unit.Format(48),
		},
		m.bp.Up(breakpoint.SM): goss.CSS{
			goss.MinHeight: unit.Format(64),
		},
	}
	return m
}

func (m Mixins) Gutters(o goss.CSS) goss.CSS {
	o[goss.PaddingLeft] = unit.Format(m.sp.Unit * 2)
	o[goss.PaddingRight] = unit.Format(m.sp.Unit * 2)
	o[m.bp.Up(breakpoint.SM)] = goss.CSS{
		goss.PaddingLeft:  unit.Format(m.sp.Unit * 3),
		goss.PaddingRight: unit.Format(m.sp.Unit * 3),
	}
	return o
}
