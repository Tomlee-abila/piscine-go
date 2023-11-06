package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

// func printInt(i int) {
// 	for _, r := range s {
// 		z01.PrintRune(r)
// 	}
// }

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	s1 := "x = 42"
	s2 := ", y = 21"
	setPoint(points)
	printStr(s1)
	printStr(s2)
	z01.PrintRune('\n')
	// printStr(s2)
	// // result += string((points.x/10)+'0') + string((points.x%10)+'0') + ", y = " + string((points.y/10)+'0') + string((points.y%10)+'0')
	// // printStr(result)
	// z01.PrintRune('x')
	// z01.PrintRune(' ')
	// z01.PrintRune('=')
	// z01.PrintRune(' ')
	// fmt.Print(((points.x / 10) + '0'))
	// z01.PrintRune(rune((points.x % 10) + '0'))
	// z01.PrintRune(',')
	// z01.PrintRune(' ')
	// z01.PrintRune('y')
	// z01.PrintRune(' ')
	// z01.PrintRune('=')
	// z01.PrintRune(' ')
	// z01.PrintRune(rune((points.x / 10) + '0'))
	// z01.PrintRune(rune((points.x % 10) + '0'))
	// z01.PrintRune('\n')
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
