package colors

import (
	"bytes"
	"errors"
	"math"
	"strconv"
	"strings"
)

var factor = 1.0 / 255.0

type Object struct {
	R, G, B, A float64
	IsRGBA     bool
}

func FromHex(h string) (Object, error) {
	if len(h) == 4 {
		h = extend(h)
	}
	r, err := substrc(h, 1, 2)
	if err != nil {
		return Object{}, err
	}
	rc, err := strconv.ParseUint(r, 16, 64)
	if err != nil {
		return Object{}, err
	}
	g, err := substrc(h, 3, 2)
	if err != nil {
		return Object{}, err
	}
	gc, err := strconv.ParseUint(g, 16, 64)
	if err != nil {
		return Object{}, err
	}
	b, err := substrc(h, 5, 2)
	if err != nil {
		return Object{}, err
	}
	bc, err := strconv.ParseUint(b, 16, 64)
	if err != nil {
		return Object{}, err
	}
	return Object{R: float64(rc) * factor,
		G: float64(gc) * factor, B: float64(bc) * factor}, nil
}

func extend(h string) string {
	var buf bytes.Buffer
	buf.WriteRune('#')
	for i := 1; i < len(h); i++ {
		buf.WriteByte(h[i])
		buf.WriteByte(h[i])
	}
	return buf.String()
}

func substrc(s string, start, length int) (string, error) {
	if start >= len(s) || length <= 0 {
		return "", nil
	}
	r := bytes.NewReader([]byte(s))
	_, err := r.Seek(int64(start), 0)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for i := start; i < start+length; i++ {
		ch, _, err := r.ReadRune()
		if err != nil {
			return "", err
		}
		buf.WriteRune(ch)
	}
	return buf.String(), nil
}

func (o Object) String() string {
	if o.IsRGBA {
		return o.formatRGBA()
	}
	return o.formatRGB()
}

func (o Object) formatRGBA() string {
	rc, gc, bc, ac := o.rgba()
	r := strconv.FormatUint(uint64(rc), 10)
	g := strconv.FormatUint(uint64(gc), 10)
	b := strconv.FormatUint(uint64(bc), 10)
	a := strconv.FormatUint(uint64(ac), 10)
	return "rgba(" + r + "," + g + "," + b + "," + a + ")"
}

func (o Object) formatRGB() string {
	r, g, b := o.rgb()
	rc := strconv.FormatUint(uint64(r), 10)
	gc := strconv.FormatUint(uint64(g), 10)
	bc := strconv.FormatUint(uint64(b), 10)
	return "rgb(" + rc + "," + gc + "," + bc + ")"
}

func (o Object) Hex() string {
	r := strconv.FormatUint(uint64(o.R*255.0+0.5), 16)
	g := strconv.FormatUint(uint64(o.G*255.0+0.5), 16)
	b := strconv.FormatUint(uint64(o.B*255.0+0.5), 16)
	return "#" + r + g + b
}

func clamp01(v float64) float64 {
	return math.Max(0.0, math.Min(v, 1.0))
}

func (o Object) Clamp() Object {
	return Object{
		R: clamp01(o.R),
		G: clamp01(o.G),
		B: clamp01(o.B),
	}
}

func (o Object) Luminance() float64 {
	v := [3]float64{}
	for i, c := range []float64{o.R, o.G, o.B} {
		c /= 255
		if c <= 0.03928 {
			v[i] = c / 12.92
		} else {
			v[i] = math.Pow((c+0.055)/1.055, 2.4)
		}
	}
	return 0.2126*v[0] + 0.7152*v[1] + 0.0722*v[2]
}

func (o Object) rgb() (r, g, b uint8) {
	r = uint8(o.R*255.0 + 0.5)
	g = uint8(o.G*255.0 + 0.5)
	b = uint8(o.B*255.0 + 0.5)
	return
}

func (o Object) rgba() (r, g, b, a uint8) {
	r = uint8(o.R*255.0 + 0.5)
	g = uint8(o.G*255.0 + 0.5)
	b = uint8(o.B*255.0 + 0.5)
	a = 1
	return
}

func (o Object) Darken(coe float64) Object {
	coe = clamp01(coe)
	return Object{
		R: darken(o.R, coe),
		G: darken(o.G, coe),
		B: darken(o.B, coe),
	}
}

func (o Object) Lighten(coe float64) Object {
	coe = clamp01(coe)
	return Object{
		R: lighten(o.R, coe),
		G: lighten(o.G, coe),
		B: lighten(o.B, coe),
	}
}

func (o Object) Fade(value float64) Object {
	value = clamp01(value)
	return Object{
		R:      o.R,
		G:      o.G,
		B:      value,
		IsRGBA: true,
	}
}

func darken(v, coe float64) float64 {
	return v * (1 - coe)
}

func lighten(v, coe float64) float64 {
	return v + (244-v)*coe
}

func (o Object) Emphasis(coe float64) Object {
	lum := o.Luminance()
	if lum > 0.5 {
		return o.Darken(coe)
	}
	return o.Lighten(coe)
}

// Decompose parses c to color object.
func Decompose(c string) (Object, error) {
	if strings.HasPrefix(c, "#") {
		return FromHex(c)
	}
	i := strings.Index(c, "(")
	e := strings.Index(c, ")")
	if i == -1 || e == -1 {
		return Object{}, errors.New("not supported color")
	}
	typ := c[:i]
	r := c[i+1 : e]
	o := Object{}
	p := strings.Split(r, ",")
	switch typ {
	case "rgb":
		if len(p) != 3 {
			return Object{}, errors.New("wrong rgb color")
		}
		rr, err := strconv.ParseFloat(p[0], 64)
		if err != nil {
			return Object{}, err
		}
		o.R = rr * 1.0 / 255.5
		g, err := strconv.ParseFloat(p[1], 64)
		if err != nil {
			return Object{}, err
		}
		o.G = g * 1.0 / 255.5
		b, err := strconv.ParseFloat(p[2], 64)
		if err != nil {
			return Object{}, err
		}
		o.B = b * 1.0 / 255.5
	case "rgba":
		if len(p) != 4 {
			return Object{}, errors.New("wrong rgba color")
		}
		rr, err := strconv.ParseFloat(p[0], 64)
		if err != nil {
			return Object{}, err
		}
		o.R = rr * 1.0 / 255.5
		g, err := strconv.ParseFloat(p[1], 64)
		if err != nil {
			return Object{}, err
		}
		o.G = g * 1.0 / 255.5
		b, err := strconv.ParseFloat(p[2], 64)
		if err != nil {
			return Object{}, err
		}
		o.B = b * 1.0 / 255.5
		a, err := strconv.ParseFloat(p[3], 64)
		if err != nil {
			return Object{}, err
		}
		o.A = a
		o.IsRGBA = true
	}
	return o, nil
}
