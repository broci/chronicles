package transition

import "testing"

func TestTransitions_Create(t *testing.T) {
	s := New()
	g := s.Create(nil)
	e := "all 300ms cubic-bezier(0.4, 0, 0.2, 1) 0ms"
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}

	g = s.Create([]string{"color"})
	e = "color 300ms cubic-bezier(0.4, 0, 0.2, 1) 0ms"
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}

	g = s.Create([]string{"color", "size"})
	e = "color 300ms cubic-bezier(0.4, 0, 0.2, 1) 0ms,size 300ms cubic-bezier(0.4, 0, 0.2, 1) 0ms"
	if g != e {
		t.Errorf("expected %s got %s", e, g)
	}
}
