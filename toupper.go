package main

import (
	"fmt"
)

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(s string) string {
	runeS := []rune(s)
	for i, str := range s {
		if lower(str) {
			runeS[i] -= 32
		}
	}
	return string(runeS)
}

func lower(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
