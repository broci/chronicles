package button

import (
	"fmt"

	"github.com/gernest/chronicles/colors"
	"github.com/gernest/chronicles/styles/transition"

	"github.com/gernest/chronicles/styles/theme"
	"github.com/gernest/chronicles/ui/component"

	css "github.com/gernest/goss"
	"honnef.co/go/js/dom"
)

// button specific classes
const (
	FlatPrimary  = "flatPrimary"
	FlatAccent   = "flatAccent"
	FlatContrast = "flatContrast"
	ColorInherit = "colorInherit"
)

// Style returns button specific css styles
func Style(t theme.Theme) css.CSS {
	tr := t.Transitions.Create([]string{
		css.BackgroundColor, css.BoxShadow,
	}, transition.Options{
		Duration: t.Transitions.Duration.Short,
	})
	contrastText := t.Palette.GetContrastText(t.Palette.Primary[colors.C500])
	raised := t.Palette.GetContrastText(t.Palette.Grey[colors.C300])
	return css.CSS{
		"root": css.CSS{
			css.LineHeight:   "1.4em",
			css.BoxSizing:    "border-box",
			css.MinWidth:     88,
			css.MinHeight:    36,
			css.BorderRadius: 2,
			css.Padding:      fmt.Sprintf("%dpx %dpx", t.Spacing.Unit, t.Spacing.Unit*2),
			css.Color:        t.Palette.Text.Primary,
			css.Transition:   tr,
			"{{.root}}:hover": css.CSS{
				css.TextDecoration: "none",
				"{{.root}} {{.disabled}}": css.CSS{
					css.Background: "transparent",
				},
			},
		},
		"@media (hover: none)": css.CSS{
			"{{.root}}:hover": css.CSS{
				css.Background: "transparent",
			},
			fmt.Sprintf("{{.%s}}:hover", FlatPrimary): css.CSS{
				css.Background: "transparent",
			},
			fmt.Sprintf("{{.%s}}:hover", FlatAccent): css.CSS{
				css.Background: "transparent",
			},
			fmt.Sprintf("{{.%s}}:hover", FlatContrast): css.CSS{
				css.Background: "transparent",
			},
		},
		"dense": css.CSS{
			css.Padding:   fmt.Sprintf("%dpx %dpx", t.Spacing.Unit-1, t.Spacing.Unit),
			css.MinWidth:  64,
			css.MinHeight: 32,
			css.FontSize:  t.Typography.PxToRem(t.Typography.FontSize.(int) - 1),
		},
		"label": css.CSS{
			css.Width:          "100%",
			css.Display:        "inherit",
			css.AlignItems:     "inherit",
			css.JustifyContent: "inherit",
		},
		FlatPrimary: css.CSS{
			css.Color: t.Palette.Primary[colors.C500],
			fmt.Sprintf("{{.%s}}:hover", FlatPrimary): css.CSS{
				css.BackgroundColor: colors.Fade(t.Palette.Primary[colors.C500], 0.12),
			},
		},
		FlatAccent: css.CSS{
			css.Color: t.Palette.Secondary[colors.A200],
			fmt.Sprintf("{{.%s}}:hover", FlatAccent): css.CSS{
				css.BackgroundColor: colors.Fade(t.Palette.Secondary[colors.A200], 0.12),
			},
		},
		FlatContrast: css.CSS{
			css.Color: contrastText,
			fmt.Sprintf("{{.%s}}:hover", FlatAccent): css.CSS{
				css.BackgroundColor: colors.Fade(contrastText, 0.12),
			},
		},
		ColorInherit: css.CSS{
			"color": "inherit",
		},
		"raised": css.CSS{
			css.Color: raised,
			css.BackgroundColor: t.Palette.Grey[colors.C300],
			css.BoxShadow: t.Shadows[2],
			"{{.raised}}{{.keyboardFocused}}": css.CSS{
				css.BoxShadow: t.Shadows[6],
			},
			"{{.raised}}:active": css.CSS{
				css.BoxShadow: t.Shadows[8],
			},
		}
	}
}

// Base is the BaseButton component.
type Base struct {
	Type                     string
	CenterRipple             bool
	DisableRipple            bool
	FocusRipple              bool
	KeyboardFocusedClassName string
	Node                     dom.Element
	Style                    css.CSS
	TabIndex                 int

	//events handlers

	OnBlur          func(dom.Event)
	OnClick         func(dom.Event)
	OnFocused       func(dom.Event)
	OnKeyBoardFocus func(dom.Event)
	OnKeyDown       func(dom.Event)
	OnKeyUp         func(dom.Event)
	OnMouseDown     func(dom.Event)
	OnMouseLeave    func(dom.Event)
	OnMouseUp       func(dom.Event)
	OnTouchEnd      func(dom.Event)
	OnTouchMove     func(dom.Event)
	OnTouchStart    func(dom.Event)
}

func (b *Base) Template() string {
	return `{{cname .classes.root| attr "class" |open .type }}{{.children}}{{close .type}}`
}

func (b *Base) Props() component.Props {
	typ := b.Type
	if typ == "" {
		typ = "button"
	}
	return component.Props{
		"type": typ,
	}
}

func (b *Base) ComponentDidMount(ctx *component.Context) error {
	b.Node = ctx.Element
	return nil
}

func (b *Base) Init(ctx *component.Context) component.Component {
	c := *b
	return &c
}

func (b *Base) ComponentStyle(t theme.Theme) css.CSS {
	return Style(t)
}


