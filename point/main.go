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
	z01.PrintRune('\n')
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	result := "x = "
	setPoint(points)
	result += string((points.x/10)+'0') + string((points.x%10)+'0') + ", y = " + string((points.y/10)+'0') + string((points.y%10)+'0')
	printStr(result)
	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
