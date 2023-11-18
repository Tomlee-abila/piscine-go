package piscine

func AlphaCount(str string) int {
	count := 0
	for _, s := range str {
		if alphabet(s) {
			count++
		}
	}
	return count
}

func alphabet(alpha rune) bool {
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
	return false
}
