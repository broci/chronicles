package button

import (
	"github.com/broci/chronicles/ui/component"
	"github.com/broci/chronicles/util"
	"github.com/broci/goss"
	css "github.com/broci/goss"
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
			"&:hover": css.CSS{
				css.TextDecoration: "none",
				"@media (hover: none)": css.CSS{
					css.Background: "transparent",
				},
				"&$disabled": css.CSS{
					css.Background: "transparent",
				},
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

type Base struct {
	Children                 component.Component
	Type                     string
	CenterRipple             bool
	DisableRipple            bool
	FocusRipple              bool
	KeyboardFocusedClassName string
	Node                     dom.Element
	Style                    *goss.Sheet
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
	return `{{open .type .class}}{{- .children -}} {{close .type}}`
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
		"class": util.Class(
			b.Style.Class["root"],
		),
	}
}

func (b *Base) ComponentDidMount(ctx *component.Context) error {
	b.Node = ctx.Element
	return nil
}

func (b *Base) Init(ctx *component.Context) component.Component {
	c := *b
	if ctx.StyleSheet != nil {
		s := ctx.StyleSheet.NewSheet()
		err := s.Parse(Style())
		if err != nil {
			panic(err)
		}
		c.Style = s
	}
	return &c
}
