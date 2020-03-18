package roman

import "strings"

type Numeral struct {
	symbol string
	value  int
}

var symbols = []Numeral{
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
