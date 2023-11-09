package piscine

func JumpOver(str string) string {
	result := ""
	for i, ch := range str {
		if (i+1)%3 == 0 {
			result += string(ch)
		}
	}
	return result
}
