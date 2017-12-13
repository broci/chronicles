package funcs

import (
	"html/template"
	"strings"

	"github.com/broci/classnames"
)

func New() template.FuncMap {
	return template.FuncMap{
		"classNames": classnames.Join,
		"open": func(src ...string) template.HTML {
			return template.HTML("<" + strings.Join(src, " ") + ">")
		},
		"close": func(src ...string) template.HTML {
			return template.HTML("</" + strings.Join(src, " ") + ">")
		},
	}
}
