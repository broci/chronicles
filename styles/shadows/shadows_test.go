package shadows

import "testing"

func TestShadows(t *testing.T) {
	s := New()
	if s[0] != "none" {
		t.Errorf("expected none got %s", s[0])
	}
}
