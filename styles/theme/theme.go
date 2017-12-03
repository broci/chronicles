package theme

import (
	"github.com/broci/chronicles/styles/breakpoint"
	"github.com/broci/chronicles/styles/palette"
	"github.com/broci/chronicles/styles/transition"
	"github.com/broci/chronicles/styles/typography"
)

type Theme struct {
	Palette     palette.Palette
	Typography  typography.Typography
	Breakpoints breakpoint.Breakpoints
	Transitions transition.Transitions
}
