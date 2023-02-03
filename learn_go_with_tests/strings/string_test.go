package strings

import "testing"

func TestStringidx(t *testing.T) {
	s1 := "chicken"
	substr1 := "ken"

	expected := 4
	got := Stringidx(s1, substr1)

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func TestBookend(t *testing.T) {
	s1 := "chicken"

	expected := "cn"
	got := Bookend(s1)

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}
