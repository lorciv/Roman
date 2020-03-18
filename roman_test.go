package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := Ator(1)
	want := "I"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}