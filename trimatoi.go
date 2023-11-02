package piscine

func TrimAtoi(s string) int {
	runeString := []rune(s)
	var Array []int
	number := 0
	minusDigit := 0
	firstDigit := -1
	count := 0

	for i, str := range runeString {
		if str == '-' && firstDigit < minusDigit {
			minusDigit = i
		}

		if numeric(str) {
			count++
			if firstDigit < 0 {
				firstDigit = i
			}
			Array = append(Array, int(str-'0'))
		}
	}

	for i := 0; i < count; i++ {
		number = (number * 10) + Array[i]
	}
	if firstDigit > minusDigit {
		number *= -1
	}
	return number
}
