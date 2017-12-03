package theme

import (
	"github.com/broci/chronicles/styles/breakpoint"
	"github.com/broci/chronicles/styles/mixins"
	"github.com/broci/chronicles/styles/palette"
	"github.com/broci/chronicles/styles/spacing"
	"github.com/broci/chronicles/styles/transition"
	"github.com/broci/chronicles/styles/typography"
)

type Theme struct {
	Direction   string
	Palette     palette.Palette
	Typography  typography.Typography
	Breakpoints breakpoint.Breakpoints
	Transitions transition.Transitions
	Mixins      mixins.Mixins
	Shadows     []string
	spacing     spacing.Spacing
}
