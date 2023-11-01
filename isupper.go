package piscine

func IsUpper(s string) bool {
	for _, str := range s {
		if !(upper(str)) {
			return false
		}
	}
	return true
}

func upper(alpha rune) bool {
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
