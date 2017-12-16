package button

import (
	"github.com/gernest/chronicles/ui/component"
	"github.com/gernest/goss"
	css "github.com/gernest/goss"
	"honnef.co/go/js/dom"
)

// button specific classes
const (
	FlatPrimary  = "flat-primary"
	FlatAccent   = "flat-accent"
	FlatContrast = "flat-contrast"
	ColorInherit = "color-inherit"
)

// Style returns button specific css styles
func Style() css.CSS {
	return css.CSS{
		"root": css.CSS{
			css.LineHeight:   "1.4em",
			css.BoxSizing:    "border-box",
			css.MinWidth:     88,
			css.MinHeight:    36,
			css.BorderRadius: 2,
			css.Padding:      `{{.theme.Spacing.Unit}}px {{multi .theme.Spacing.Unit  2}}px`,
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
		},
		"dense": css.CSS{
			css.MinWidth:  64,
			css.MinHeight: 32,
		},
		"label": css.CSS{
			css.Width:          "100%",
			css.Display:        "inherit",
			css.AlignItems:     "inherit",
			css.JustifyContent: "inherit",
		},
		FlatPrimary: css.CSS{},
		FlatAccent:  css.CSS{},
		ColorInherit: css.CSS{
			"color": "inherit",
		},
	}
}

// Base is the BaseButton component.
type Base struct {
	Children                 component.Component
	Type                     string
	CenterRipple             bool
	DisableRipple            bool
	FocusRipple              bool
	KeyboardFocusedClassName string
	Node                     dom.Element
	Style                    goss.CSS
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
	return `{{cname .classes.root| attr "class" |open .type }}{{- .children -}}{{close .type}}`
}

func (b *Base) Props() component.Props {
	typ := b.Type
	if typ == "" {
		typ = "button"
	}
	children := ""
	if b.Children != nil {
		children = b.Children.Template()
	}
	return component.Props{
		"type":     typ,
		"children": children,
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

func (b *Base) ComponentStyle() goss.CSS {
	s := Style()
	if b.Style != nil {
		for k, v := range b.Style {
			s[k] = v
		}
	}
	return s
}
