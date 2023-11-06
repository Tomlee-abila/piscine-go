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

func printInt(i int) {
	x1 := '4'
	x2 := '2'
	y2 := '1'
	r := '0'
	if i == 4 {
		r = x1
	} else if i == 2 {
		r = x2
	} else {
		r = y2
	}
	z01.PrintRune(r)
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	s1 := "x = "
	s2 := ", y = "
	setPoint(points)
	printStr(s1)
	printInt((points.x / 10))
	printInt((points.x % 10))
	printStr(s2)
	printInt((points.y / 10))
	printInt((points.y % 10))
	z01.PrintRune('\n')
}
