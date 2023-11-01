package piscine

func IsPrintable(s string) bool {
	for _, str := range s {
		if !(alphaNum(str)) {
			return false
		}
	}
	return true
}

func alphaNum(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := '0'; a <= '9'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
