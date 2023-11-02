package main

import (
	"fmt"
)

func main() {
	fmt.Println(ToLower("HELLLO! HOW ARE YOU?"))
}

func ToLower(s string) string {
	runeS := []rune(s)
	for i, str := range s {
		if upper(str) {
			runeS[i] += 32
		}
	}
	return string(runeS)
}

func upper(alpha rune) bool {
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
