package main

import "github.com/01-edu/z01"

func main() {
	DescendComb()
}

// package piscine

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if (int(i)*10)+int(j) > (int(k)*10)+int(l) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						// }
						// if k >= l {
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						// z01.PrintRune('\n')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
