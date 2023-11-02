package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	title := []rune(os.Args[0])
	for i := 2; i < len(title); i++ {
		z01.PrintRune(title[i])
	}
	z01.PrintRune('\n')
}
