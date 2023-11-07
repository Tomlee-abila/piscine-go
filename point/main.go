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
	setPoint(points)
	s1 := "x = "
	s2 := ", y = "
	printStr(s1)
	printInt((points.x / 10))
	printInt((points.x % 10))
	printStr(s2)
	printInt((points.y / 10))
	printInt((points.y % 10))
	z01.PrintRune('\n')
	// vziPusU.m#cc4Bi
}
func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(i int) {
	r := '0'
	for j := 0; j < i; j++ {
		r++
	}
	z01.PrintRune(r)
}
