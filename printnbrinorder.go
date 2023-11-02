package piscine

import "github.com/01-edu/z01.PrintRune"

func PrintNbrInOrder(n int) {
	count := 0
	temp := n
	fNumber := 0
	temp1 := 0
	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n))
		return
	} else if n < 0 {
		z01.PrintRune('0')
		return
	}
	for i := 0; temp > 0; i++ {
		temp /= 10
		count++
	}
	number := make([]int, count)
	temp = n
	for i := 0; temp > 0; i++ {
		number[i] = temp % 10
		temp /= 10
	}
	for j := 0; j < count; j++ {
		for i := 0; i < count-1; i++ {
			if number[i] > number[i+1] {
				temp1 = number[i]
				number[i] = number[i+1]
				number[i+1] = temp1
			}
		}
	}
	for i := 0; i < count; i++ {
		fNumber = (fNumber * 10) + number[i]
	}
	runeNumber := []rune(fNumber)
	for i := 0; i < count; i++ {
		z01.PrintRune(runeNumber[i])
	}
}
