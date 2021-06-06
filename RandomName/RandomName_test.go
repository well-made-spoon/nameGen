package nameGen

import "testing"

func TestRandomName(t *testing.T) {
	got := RandomName()
	if len(got) < 1 {
		t.Errorf("got %s; want a string", got)
	}
}
