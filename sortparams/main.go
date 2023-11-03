package main

import (
	"os"

	"github.com/01-edu/z01"
)

func insertionSort(args []string) {
	for i := 1; i < len(args); i++ {
		j := i
		for j > 0 && args[j] < args[j-1] {
			args[j], args[j-1] = args[j-1], args[j]
			j--
		}
	}
}

func main() {
	args := os.Args[1:]
	insertionSort(args)

	for _, arg := range args {
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
