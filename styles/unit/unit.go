package unit

import (
	"unicode/utf8"

	"github.com/broci/chronicles/util"
)

type Unit interface{}

func Px(u Unit) string {
	if s, ok := u.(string); ok {
		if !isPx(s) {
			panic("trying to convert non px string unit: " + s)
		}
		return s
	}
	return toString(u) + "px"
}

func Rem(u Unit) string {
	if s, ok := u.(string); ok {
		if !isPx(s) {
			panic("trying to convert non rem string unit: " + s)
		}
		return s
	}
	return toString(u) + "rem"
}

func toString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case int:
		return util.FormatInt(t)
	case int64:
		return util.FormatInt64(t)
	default:
		panic("Unkown unit value")
	}
}

func isPx(s string) bool {
	if len(s) < 3 {
		return false
	}
	x, size := utf8.DecodeLastRuneInString(s)
	if x != 'x' {
		return false
	}
	p, _ := utf8.DecodeLastRuneInString(s[:len(s)-size])
	return p == 'p'
}

func isRem(s string) bool {
	if len(s) < 4 {
		return false
	}
	m, size := utf8.DecodeLastRuneInString(s)
	if m != 'm' {
		return false
	}
	e, size := utf8.DecodeLastRuneInString(s[:len(s)-size])
	if e != 'e' {
		return false
	}
	r, _ := utf8.DecodeLastRuneInString(s[:len(s)-size])
	return r == 'r'
}

func Em(u Unit) string {
	if f, ok := u.(float64); ok {
		util.FormatFloat(f)
	}
	return toString(u) + "em"
}

func Format(u Unit) string {
	if s, ok := u.(string); ok {
		return s
	}
	return Px(u)
}
