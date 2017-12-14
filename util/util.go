package util

import (
	"fmt"
	"strconv"

	"github.com/gernest/classnames"
)

func FormatInt(v int) string {
	return FormatInt64(int64(v))
}

func FormatInt64(v int64) string {
	return strconv.FormatInt(v, 10)
}

func FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func Class(n ...interface{}) string {
	v := classnames.Join(n...)
	if v != "" {
		return fmt.Sprintf(`class="%s"`, v)
	}
	return ""
}
