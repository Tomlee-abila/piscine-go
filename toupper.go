package piscine

func ToUpper(s string) string {
	runeS := []rune(s)
	for i, str := range s {
		if lower(str) {
			runeS[i] -= 32
		}
	}
	return string(runeS)
}
