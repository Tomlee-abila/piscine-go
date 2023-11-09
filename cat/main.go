package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if openFile(os.Args[i]) {
				readFile(os.Args[i])
			}
		}
		return
	}
}

func openFile(fName string) bool {
	file, err := os.Open(fName)
	if err != nil {
		PrintStr("ERROR: ")
		PrintStr(err.Error())
		z01.PrintRune('\n')
		os.Exit(1)
		return false
	}
	defer file.Close()
	return true
}

func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		PrintStr(err.Error())
		z01.PrintRune('\n')
		return
	}
	PrintStr(string(content))
	// filename := os.Args[1]
	// // Read the content of the file
	// content, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// // Print the content of the file as a string
	// fmt.Print(string(content))
}

func PrintStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
	// z01.PrintRune('\n')
}
