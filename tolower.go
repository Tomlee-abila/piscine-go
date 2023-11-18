package piscine

func ToLower(s string) string {
	runeS := []rune(s)
	for i, str := range s {
		if upper(str) {
			runeS[i] += 32
		}
	}
	return string(runeS)
}
