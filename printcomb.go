package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	z01.Printf("012")

	for i := 13; i <= 999; i++ {
		
		firstDigit := i / 100         
		secondDigit := (i / 10) % 10  
		thirdDigit := i % 10          

		if firstDigit < secondDigit && secondDigit < thirdDigit {

			z01.Printf(", %03d", i)
		}	
		
	}
}


