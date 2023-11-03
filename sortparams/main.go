package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		title := []rune(os.Args[i])
		for i := 0; i < len(title); i++ {
			z01.PrintRune(title[i])
		}
		z01.PrintRune('\n')
	}
	// for j := 0; j < len(os.Args)-1; j++ {
	// 	for i := 1; i < len(os.Args)-1; j++ {
	// 		if os.Args[i] > os.Args[i+1] {
	// 			minValue := os.Args[i]
	// 			os.Args[i] = os.Args[i+1]
	// 			os.Args[i+1] = minValue
	// 		}
	// 	}
	// }
}

