package colors

import "bytes"
import "strconv"

type Object struct {
	R, G, B, A uint8
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
	return Object{R: uint8(rc), G: uint8(gc), B: uint8(bc)}, nil
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
	r := strconv.FormatUint(uint64(o.R), 10)
	g := strconv.FormatUint(uint64(o.G), 10)
	b := strconv.FormatUint(uint64(o.B), 10)
	i := o.A
	if i == 0 {
		i = 1
	}
	a := strconv.FormatUint(uint64(i), 10)
	return "rgba(" + r + "," + g + "," + b + "," + a + ")"
}

func (o Object) formatRGB() string {
	r := strconv.FormatUint(uint64(o.R), 10)
	g := strconv.FormatUint(uint64(o.G), 10)
	b := strconv.FormatUint(uint64(o.B), 10)
	return "rgb(" + r + "," + g + "," + b + ")"
}

func (o Object) Hex() string {
	r := strconv.FormatUint(uint64(o.R), 16)
	g := strconv.FormatUint(uint64(o.G), 16)
	b := strconv.FormatUint(uint64(o.B), 16)
	return "#" + r + g + b
}
