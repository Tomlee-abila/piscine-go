package piscine

func ToUpper(s string) string {
	upper := []rune(s)
	for i, str := range s {
		if lower(str) {
			upper[i] -= 32
		}
	}
	return string(upper)
}
