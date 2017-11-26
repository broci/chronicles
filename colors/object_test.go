package colors

import "testing"

func TestObject(t *testing.T) {
	o, err := FromHex("#5f55f5")
	if err != nil {
		t.Fatal(err)
	}
	e := "rgb(95,85,245)"
	g := o.String()
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
	x := o.Hex()
	if x != "#5f55f5" {
		t.Errorf("expected #5f55f5 got %s", x)
	}

	o.IsRGBA = true
	e = "rgba(95,85,245,1)"
	g = o.String()
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}

	o, err = FromHex("#5f5")
	if err != nil {
		t.Fatal(err)
	}

	e = "rgb(85,255,85)"
	g = o.String()
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
}

func TestSubStr(t *testing.T) {
	sample := []struct {
		src           string
		start, length int
		exp           string
	}{
		{"#5f55f5", 1, 2, "5f"},
		{"#5f55f5", 3, 2, "55"},
		{"#5f55f5", 5, 2, "f5"},
	}
	for _, e := range sample {
		g, err := substrc(e.src, e.start, e.length)
		if err != nil {
			t.Fatal(err)
		}
		if g != e.exp {
			t.Errorf("expected %s got %s", e.exp, g)
		}
	}
}
