package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for x := 0; x < 8; x++ {
		for y := x + 1; y < 9; y++ {
			for k := y + 1; k < 10; k++ {
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + y))
				z01.PrintRune(rune('0' + k))
				if x != 7 || y != 8 || k != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
