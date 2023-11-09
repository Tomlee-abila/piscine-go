package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			if openFile(os.Args[i]) {
				readFile(os.Args[i])
			}
		}
		return
	} else if len(os.Args) < 2 {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello")
			i = 0
		}
	}
}

func openFile(fName string) bool {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return false
	}
	defer file.Close()
	return true
}

func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(content))
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
