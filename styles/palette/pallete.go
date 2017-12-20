package palette

import "github.com/gernest/chronicles/colors"

type ShadeText struct {
	Primary      string
	Secondary    string
	Disabled     string
	Hint         string
	Icon         string
	Divider      string
	LightDivider string
}

type ShadeInput struct {
	BottomLine string
	HelperText string
	LabelText  string
	InputText  string
	Disabled   string
}
type ShadeAction struct {
	Active   string
	Disabled string
}

type ShadeBackground struct {
	Default      string
	Paper        string
	AppBar       string
	ContentFrame string
	Sstatus      string
}

type Shade struct {
	Text       ShadeText
	Input      ShadeInput
	Action     ShadeAction
	Background ShadeBackground
}

type Palette struct {
	Common    colors.CommonColors
	Type      colors.Contrast
	Primary   colors.Color
	Secondary colors.Color
	Error     colors.Color
	Grey      colors.Color
	Shades    struct {
		dark  Shade
		light Shade
	}
	Text       ShadeText
	Input      ShadeInput
	Action     ShadeAction
	Background ShadeBackground
}

var grey = colors.Grey()
var common = colors.Common()

// Light is light shade.
func Light() Shade {
	return Shade{
		Text: ShadeText{
			Primary:      "rgba(0, 0, 0, 0.87)",
			Secondary:    "rgba(0, 0, 0, 0.54)",
			Disabled:     "rgba(0, 0, 0, 0.38)",
			Hint:         "rgba(0, 0, 0, 0.38)",
			Icon:         "rgba(0, 0, 0, 0.38)",
			Divider:      "rgba(0, 0, 0, 0.12)",
			LightDivider: "rgba(0, 0, 0, 0.075)",
		},
		Input: ShadeInput{
			BottomLine: "rgba(0, 0, 0, 0.42)",
			HelperText: "rgba(0, 0, 0, 0.54)",
			LabelText:  "rgba(0, 0, 0, 0.54)",
			InputText:  "rgba(0, 0, 0, 0.87)",
			Disabled:   "rgba(0, 0, 0, 0.42)",
		},
		Action: ShadeAction{
			Active:   "rgba(0, 0, 0, 0.54)",
			Disabled: "rgba(0, 0, 0, 0.26)",
		},
		Background: ShadeBackground{
			Default:      grey[colors.C50],
			Paper:        common.White,
			AppBar:       grey[colors.C100],
			ContentFrame: grey[colors.C200],
		},
	}
}

func Dark() Shade {
	return Shade{
		Text: ShadeText{
			Primary:      "rgba(255, 255, 255, 1)",
			Secondary:    "rgba(255, 255, 255, 0.7)",
			Disabled:     "rgba(255, 255, 255, 0.5)",
			Hint:         "rgba(255, 255, 255, 0.5)",
			Icon:         "rgba(255, 255, 255, 0.5)",
			Divider:      "rgba(255, 255, 255, 0.12)",
			LightDivider: "rgba(255, 255, 255, 0.075)",
		},
		Input: ShadeInput{
			BottomLine: "rgba(255, 255, 255, 0.7)",
			HelperText: "rgba(255, 255, 255, 0.7)",
			LabelText:  "rgba(255, 255, 255, 0.7)",
			InputText:  "rgba(255, 255, 255, 1)",
			Disabled:   "rgba(255, 255, 255, 0.5)",
		},
		Action: ShadeAction{
			Active:   "rgba(255, 255, 255, 1)",
			Disabled: "rgba(255, 255, 255, 0.3)",
		},
		Background: ShadeBackground{
			Default:      "#303030",
			Paper:        grey[colors.C800],
			AppBar:       grey[colors.C900],
			ContentFrame: grey[colors.C900],
		},
	}
}

func New(c colors.Contrast) Palette {
	lightShade := Light()
	darkShade := Dark()
	p := Palette{
		Common:    colors.Common(),
		Type:      c,
		Primary:   colors.Indigo(),
		Secondary: colors.Pink(),
		Error:     colors.Red(),
	}
	p.Shades.light = lightShade
	p.Shades.dark = darkShade
	switch c {
	case colors.LightContrast:
		p.Text = lightShade.Text
		p.Input = lightShade.Input
		p.Action = lightShade.Action
		p.Background = lightShade.Background
	case colors.DarkContrast:
		p.Text = darkShade.Text
		p.Input = darkShade.Input
		p.Action = darkShade.Action
		p.Background = darkShade.Background
	default:
		panic("unknown contrast value : " + c.Value())
	}
	return p
}

// GetContrastText calculate contrast for text.
func (p Palette) GetContrastText(hue string) string {
	if colors.GetContrastRatio(hue, p.Common.Black) < 7 {
		return p.Shades.dark.Text.Primary
	}
	return p.Shades.light.Text.Primary
}
