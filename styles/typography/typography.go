package typography

import (
	"strconv"

	"github.com/broci/chronicles/styles/palette"
	"github.com/broci/chronicles/styles/unit"
)

type TextStyle uint32

const (
	Button TextStyle = iota
	Display1
	Display2
	Display3
	Display4
	Headline
	Title
	Subheading
	Body1
	Body2
	Caption
)

type Style struct {
	Color         string
	FontSize      unit.Unit
	FontFamily    string
	FontWeight    unit.Unit
	LetterSpacing string
	MarginLeft    string
	LineHeight    unit.Unit
	TextTransform string
}

type Typography struct {
	FontFamily        string
	FontSize          unit.Unit
	FontWeightLight   unit.Unit
	FontWeightReqular unit.Unit
	FontWeightMedium  unit.Unit
	HTMLFontSize      unit.Unit
	Style             map[TextStyle]Style
	PxToRem           func(int) string
}

type Opts struct {
	FontSize     unit.Unit
	HTMLFontSize unit.Unit
}

func New(p *palette.Palette, opts ...Opts) *Typography {
	t := &Typography{
		FontFamily:        `"Roboto", "Helvetica", "Arial", sans-serif`,
		FontSize:          14,
		FontWeightLight:   300,
		FontWeightReqular: 400,
		FontWeightMedium:  500,
		HTMLFontSize:      16,
	}
	if len(opts) > 0 {
		o := opts[0]
		if o.FontSize != nil {
			t.FontSize = o.FontSize
		}
		if o.HTMLFontSize != nil {
			t.HTMLFontSize = o.HTMLFontSize
		}
	}
	t.PxToRem = func(v int) string {
		return PxToRem(v, t.HTMLFontSize.(int))
	}
	t.Style = map[TextStyle]Style{
		Display4: {
			FontSize:      t.PxToRem(112),
			FontWeight:    t.FontWeightLight,
			FontFamily:    t.FontFamily,
			LineHeight:    round(128.0 / 112.0),
			LetterSpacing: "-.06em",
			Color:         p.Text.Secondary,
		},
		Display3: {
			FontSize:      t.PxToRem(56),
			FontWeight:    t.FontWeightReqular,
			FontFamily:    t.FontFamily,
			LetterSpacing: "-.02em",
			LineHeight:    round(73.0 / 56.0),
			MarginLeft:    "-.04em",
			Color:         p.Text.Secondary,
		},
		Display2: {
			FontSize:      t.PxToRem(45),
			FontWeight:    t.FontWeightReqular,
			FontFamily:    t.FontFamily,
			LetterSpacing: "-.02em",
			LineHeight:    round(48.0 / 45.0),
			MarginLeft:    "-.04em",
			Color:         p.Text.Secondary,
		},
		Display1: {
			FontSize:      t.PxToRem(34),
			FontWeight:    t.FontWeightReqular,
			FontFamily:    t.FontFamily,
			LetterSpacing: "-.02em",
			LineHeight:    round(41.0 / 34.0),
			MarginLeft:    "-.04em",
			Color:         p.Text.Secondary,
		},
		Headline: {
			FontSize:   t.PxToRem(24),
			FontWeight: t.FontWeightReqular,
			FontFamily: t.FontFamily,
			LineHeight: round(32.5 / 24),
			Color:      p.Text.Primary,
		},
		Title: {
			FontSize:   t.PxToRem(21),
			FontWeight: t.FontWeightReqular,
			FontFamily: t.FontFamily,
			LineHeight: round(24.0 / 16),
			Color:      p.Text.Primary,
		},
		Body2: {
			FontSize:   t.PxToRem(14),
			FontWeight: t.FontWeightMedium,
			FontFamily: t.FontFamily,
			LineHeight: round(24.0 / 14),
			Color:      p.Text.Primary,
		},
		Body1: {
			FontSize:   t.PxToRem(14),
			FontWeight: t.FontWeightReqular,
			FontFamily: t.FontFamily,
			LineHeight: round(20.4 / 14),
			Color:      p.Text.Primary,
		},
		Caption: {
			FontSize:   t.PxToRem(12),
			FontWeight: t.FontWeightReqular,
			FontFamily: t.FontFamily,
			LineHeight: round(16.5 / 12),
			Color:      p.Text.Secondary,
		},
		Button: {
			FontSize:      t.PxToRem(t.FontSize.(int)),
			FontWeight:    t.FontWeightReqular,
			TextTransform: "uppercase",
			FontFamily:    t.FontFamily,
		},
	}
	return t
}

func PxToRem(v int, htmlFontSize int) string {
	return strconv.FormatInt(int64(v/htmlFontSize), 10) + "rem"
}

func Px(v int) string {
	return strconv.FormatInt(int64(v), 10) + "px"
}

func round(value float64) float64 {
	return (value * 1e5) / 1e5
}
