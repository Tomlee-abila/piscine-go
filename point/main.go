package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	// result := "x = "
	setPoint(points)
	// result += string((points.x/10)+'0') + string((points.x%10)+'0') + ", y = " + string((points.y/10)+'0') + string((points.y%10)+'0')
	// printStr(result)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune((points.x / 10) + '0'))
	z01.PrintRune(rune((points.x % 10) + '0'))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune(rune((points.x / 10) + '0'))
	z01.PrintRune(rune((points.x % 10) + '0'))
	z01.PrintRune('\n')
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
