package theme

import (
	"github.com/gernest/chronicles/colors"
	"github.com/gernest/chronicles/styles/breakpoint"
	"github.com/gernest/chronicles/styles/mixins"
	"github.com/gernest/chronicles/styles/palette"
	"github.com/gernest/chronicles/styles/shadows"
	"github.com/gernest/chronicles/styles/spacing"
	"github.com/gernest/chronicles/styles/transition"
	"github.com/gernest/chronicles/styles/typography"
	"github.com/gernest/chronicles/styles/zindex"
)

type Theme struct {
	Direction   string
	Palette     palette.Palette
	Typography  typography.Typography
	Breakpoints breakpoint.Breakpoints
	Transitions transition.Transitions
	Mixins      mixins.Mixins
	Shadows     []string
	Spacing     spacing.Spacing
	Zindex      zindex.Zindex
}

func New(contrast colors.Contrast) Theme {
	t := Theme{
		Direction:   "ltr",
		Palette:     palette.New(contrast),
		Breakpoints: breakpoint.New(),
		Transitions: transition.New(),
		Shadows:     shadows.New(),
		Spacing:     spacing.New(),
		Zindex:      zindex.New(),
	}
	t.Mixins = mixins.New(t.Breakpoints, t.Spacing)
	t.Typography = typography.New(t.Palette)
	return t
}
