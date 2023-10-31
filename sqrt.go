package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	sqr := 0
	for i := 1; i <= nb; i++ {
		sqr = i * i
		if sqr == nb {
			return sqr
		} else if i == nb {
			return 0
		}
	}
}
