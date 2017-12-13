package funcs

import "html/template"
import "github.com/broci/classnames"

func New() template.FuncMap {
	return template.FuncMap{
		"classNames": classnames.Join,
		"open": func(src string) template.HTML {
			return template.HTML("<" + src + ">")
		},
		"close": func(src string) template.HTML {
			return template.HTML("</" + src + ">")
		},
	}
}
