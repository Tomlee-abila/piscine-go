package piscine

func IsNumeric(s string) bool {
	for _, str := range s {
		if !(numeric(str)) {
			return false
		}
	}
	return true
}

func numeric(alpha rune) bool {
	for a := '0'; a <= '9'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
