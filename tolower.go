package piscine

func IsLower(s string) string {
	runeS := []rune(s)
	for i, str := range s {
		if upper(str) {
			runeS[i] += 32
		}
	}
	return string(runeS)
}
