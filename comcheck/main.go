package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	isTrue := false
	for i := 0; i < len(args); i++ {
		if args[i] == "01" || args[i] == "galaxy" || args[i] == "galaxy 01" {
			isTrue = true
		}
	}
	if isTrue {
		fmt.Println("Alert!!!")
	}
}
