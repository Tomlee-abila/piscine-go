package piscine

func NRune(s string, n int) rune {
	number := len(s)
	sentence := []rune(s)
	if n > number || n < 0 {
		return 0
	} else {
		return sentence[n-1]
	}
}
