package funcs

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gernest/classnames"
)

func New() template.FuncMap {
	return template.FuncMap{
		"cname": classnames.Join,
		"open": func(src ...string) template.HTML {
			return template.HTML("<" + strings.Join(src, " ") + ">")
		},
		"close": func(src ...string) template.HTML {
			return template.HTML("</" + strings.Join(src, " ") + ">")
		},
		"attr": func(key string, value string) string {
			return fmt.Sprintf(`%s="%s"`, key, value)
		},
		"multi": func(a, b int) int {
			return a * b
		},
	}
}
