package piscine

func PrintNbrInOrder(n int) {
	count := 0
	temp := n
	fNumber := 0
	temp1 := 0
	if n < 10 && n >= 0 {
		fmt.Print(n)
		return
	} else if n < 0 {
		fmt.Print("0")
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
	fmt.Print(fNumber)
}
