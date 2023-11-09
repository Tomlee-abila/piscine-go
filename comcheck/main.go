package main

import (
	"fmt"
	"os"
)

func main() {
	alert := false
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if compare(os.Args[i]) {
				alert = true
			}
		}
	}
	if alert {
		fmt.Println("Alert!!!")
	}
}

func compare(s string) bool {
	if len(s) > 1 {

		arr := []rune(s)
		for i := 0; i < len(arr); i++ {
			if arr[i] == '0' {
				if arr[i] == '0' && arr[i+1] == '1' {
					return true
				}
			} else if arr[i] == 'g' {
				if arr[i] == 'g' && arr[i+1] == 'a' && arr[i+2] == 'l' && arr[i+3] == 'a' && arr[i+4] == 'x' && arr[i+5] == 'y' {
					return true
				}
			}
		}
	}
	return false
}
