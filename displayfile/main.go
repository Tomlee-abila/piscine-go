package main

import (
	"os"

	"fmt"
	"io/ioutil"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	filename := os.Args[1]
	// Read the content of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	// Print the content of the file as a string
	fmt.Print(string(content))

}
