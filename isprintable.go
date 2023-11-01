package piscine

func IsPrintable(s string) bool {
	for _, str := range s {
		if charIsPrintable(str) {
			return false
		}
	}
	return true
}

func charIsPrintable(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
