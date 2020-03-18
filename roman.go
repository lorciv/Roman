package roman

import "strings"

type symbol struct {
	symbol string
	value  int
}

var symbols = []symbol{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// Ator converts Arabic numerals to Roman numerals.
func Ator(arabic int) string {
	res := strings.Builder{}
	for _, sym := range symbols {
		for arabic >= sym.value {
			res.WriteString(sym.symbol)
			arabic -= sym.value
		}
	}
	return res.String()
}

// Rtoa converts Roman numerals to Arabic numerals.
func Rtoa(roman string) int {
	res := 0
	for _, sym := range symbols {
		for strings.HasPrefix(roman, sym.symbol) {
			res += sym.value
			roman = roman[len(sym.symbol):]
		}
	}
	return res
}
