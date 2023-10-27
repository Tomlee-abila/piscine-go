package piscine

import (
	"fmt"
)

func PrintStr(s string) {
	aString := []byte(s)

	for i := 0; i < len(aString); i++ {
		fmt.Printf("%c", aString[i])
	}
	fmt.Println()
}
