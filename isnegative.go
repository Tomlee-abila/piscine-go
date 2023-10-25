package main

import "github.com/01-edu/z01"

func isNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}

func main() {
	isNegative(4)
	isNegative(6)
	isNegative(-3)
}