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
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{16, "XVI"},
		{17, "XVII"},
		{18, "XVIII"},
		{19, "XIX"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
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
