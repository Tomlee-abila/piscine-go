package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		title := []rune(os.Args[i])
		for i := 0; i < len(title); i++ {
			z01.PrintRune(title[i])
		}
		z01.PrintRune('\n')
	}
}
