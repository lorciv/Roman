package roman

import "testing"

func TestRomanNumerals(t *testing.T) {
	t.Run("convert 1", func(t *testing.T) {
		got := Ator(1)
		want := "I"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("convert 2", func(t *testing.T) {
		got := Ator(2)
		want := "II"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
