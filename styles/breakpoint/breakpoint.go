package breakpoint

import (
	"strconv"
)

// Media is a breakpoint in a media query.
type Media uint

// supported breakpoints
const (
	XS Media = iota + 1
	SM
	MD
	LG
	XL
)

func (m Media) String() string {
	switch m {
	case XS:
		return "xs"
	case SM:
		return "sm"
	case MD:
		return "md"
	case LG:
		return "lg"
	case XL:
		return "xl"
	default:
		return "md"
	}
}

// Value returns breakpoint value.
func (m Media) Value() int64 {
	switch m {
	case XS:
		return 0
	case SM:
		return 600
	case MD:
		return 960
	case LG:
		return 1280
	case XL:
		return 1920
	default:
		return 960
	}
}

// ValueString returns string representation of breakpoint value. This helps to
// avoid calling strconv.
func (m Media) ValueString() string {
	switch m {
	case XS:
		return "0"
	case SM:
		return "600"
	case MD:
		return "960"
	case LG:
		return "1280"
	case XL:
		return "1920"
	default:
		return "960"
	}
}

// Up sets breakpoint with min-width.
func Up(m Media, unit string) string {
	return "@media (min-width:" + m.ValueString() + unit + ")"
}

// Down set max-width.
func Down(m Media, unit string) string {
	v := strconv.FormatFloat(float64(m.Value())-0.05, 'f', 2, 64)
	return "@media (max-width:" + v + unit + ")"
}

type Breakpoints interface {
	Up(Media) string
	Down(Media) string
}

type BP struct{}

func New() BP {
	return BP{}
}
func (BP) Up(m Media) string {
	return Up(m, "px")
}

func (BP) Down(m Media) string {
	return Down(m, "px")
}
