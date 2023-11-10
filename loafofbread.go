package piscine

func LoafOfBread(str string) string {
	result := ""
	count := 0
	if len(str) <= 5 && str != "" {
		result += "Invalid Output\n"
		return result
	}
	for _, ch := range str {
		if count == 5 {
			result += string(' ')
			count = 0
		} else {
			if !(ch == ' ') {
				count++
				result += string(ch)
			}
		}
	}
	result += string('\n')
	return result
}
