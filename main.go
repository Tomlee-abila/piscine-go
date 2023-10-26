package main

import "piscine"
import "github.com/01-edu/z01"


func main() {
	z01.PrintRune('0')

	for i := 13; i <= 999; i++ {

		firstDigit := char := rune(i / 100)
		secondDigit := char := rune((i / 10) % 10)
		thirdDigit := char := rune(i % 10)

		z01.PrintRune(i)
		if firstDigit < secondDigit && secondDigit < thirdDigit {

			z01.PrintRune(",")
			z01.PrintRune(firstDigit)
			z01.PrintRune(secondDigit)
			z01.PrintRune(thirdDigit)
		}

	}
}

// package main

// import "github.com/01-edu/z01"


// func main() {
// 	for i := 0; i <= 100; i++ {
// 		z01.PrintRune(rune(i))
// 	}

// }