package button

import (
	css "github.com/broci/goss"
)

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
	}
}
