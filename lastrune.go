package main

func LastRune(s string) rune {
	number := len(s)
	sentence := []rune(s)
	return sentence[number-1]
}
