package theme

import (
	"github.com/broci/chronicles/colors"
	"github.com/broci/chronicles/styles/breakpoint"
	"github.com/broci/chronicles/styles/mixins"
	"github.com/broci/chronicles/styles/palette"
	"github.com/broci/chronicles/styles/shadows"
	"github.com/broci/chronicles/styles/spacing"
	"github.com/broci/chronicles/styles/transition"
	"github.com/broci/chronicles/styles/typography"
	"github.com/broci/chronicles/styles/zindex"
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
