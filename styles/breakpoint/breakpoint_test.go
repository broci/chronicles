package breakpoint

import "testing"

func TestUp(t *testing.T) {
	g := Up(XS, "px")
	e := "@media (min-width:0px)"
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
}
func TestDown(t *testing.T) {
	g := Down(MD, "px")
	e := "@media (max-width:959.95px)"
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
}
