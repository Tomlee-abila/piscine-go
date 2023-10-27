package piscine

import (
	"fmt"
)

func PrintStr(s string) {

	aString := []byte(s)

	// aString := "Hello, World!"

	for i := 0; i < len(aString); i++ {
		fmt.Printf("%c", aString[i])
	}

	// for i, _ := range aString {
	// 	fmt.Printf("%v", aString[i])
	// }

	fmt.Println()

}
