package roman

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		arabic int
		roman  string
	}{
		{1, "I"},
		{2, "II"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("convert %d to %s", test.arabic, test.roman)
		t.Run(descr, func(t *testing.T) {
			got := Ator(test.arabic)
			if got != test.roman {
				t.Errorf("got %q, want %q", got, test.roman)
			}
		})
	}
}
