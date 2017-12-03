package transition

import "time"

type Easing struct {
	InOut string
	Out   string
	In    string
	Sharp string
}

type Duration struct {
	Shortest       time.Duration
	Shorter        time.Duration
	Short          time.Duration
	Standard       time.Duration
	Complex        time.Duration
	EnteringScreen time.Duration
	LeavingScreen  time.Duration
}

type Options struct {
	Duration time.Duration
	Easing   string
	Delay    time.Duration
}

type Transitions struct {
	Easing   Easing
	Duration Duration
}

func (t Transitions) Create(props []string, opts ...Options) string {
	dur := t.Duration.Standard
	ease := t.Easing.InOut
	var delay time.Duration
	if len(opts) > 0 {
		o := opts[0]
		if o.Duration != 0 {
			dur = o.Duration
		}
		if o.Easing != "" {
			ease = o.Easing
		}
		delay = o.Delay
	}
	if props == nil {
		props = []string{"all"}
	}
	return join(",", format(props, dur, delay, ease)...)
}

func format(props []string, duration, delay time.Duration, ease string) []string {
	var o []string
	for _, v := range props {
		d := "0ms"
		if delay != 0 {
			d = delay.String()
		}
		o = append(o, join(" ", v, duration.String(), ease, d))
	}
	return o
}

func join(sep string, v ...string) string {
	o := ""
	for i, s := range v {
		if i == 0 {
			o += s
		} else {
			o += sep + s
		}
	}
	return o
}

func DefaultEasing() Easing {
	return Easing{
		// This is the most common easing curve.
		InOut: "cubic-bezier(0.4, 0, 0.2, 1)",
		// Objects enter the screen at full velocity from off-screen and
		// slowly decelerate to a resting point.
		Out: "cubic-bezier(0.0, 0, 0.2, 1)",
		// Objects leave the screen at full velocity. They do not decelerate when off-screen.
		In: "cubic-bezier(0.4, 0, 1, 1)",
		// The sharp curve is used by objects that may return to the screen at any time.
		Sharp: "cubic-bezier(0.4, 0, 0.6, 1)",
	}
}

func DefaultDuration() Duration {
	return Duration{
		Shortest: 150 * time.Millisecond,
		Shorter:  200 * time.Millisecond,
		Short:    250 * time.Millisecond,
		// most basic recommended timing
		Standard: 300 * time.Millisecond,
		// this is to be used in complex animations
		Complex: 375 * time.Millisecond,
		// recommended when something is entering screen
		EnteringScreen: 225 * time.Millisecond,
		// recommended when something is leaving screen
		LeavingScreen: 195 * time.Millisecond,
	}
}

func New() Transitions {
	return Transitions{
		Easing:   DefaultEasing(),
		Duration: DefaultDuration(),
	}
}
