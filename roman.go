package roman

import "strings"

func Ator(arabic int) string {
	res := strings.Builder{}
	for i := 0; i < arabic; i++ {
		res.WriteString("I")
	}
	return res.String()
}
