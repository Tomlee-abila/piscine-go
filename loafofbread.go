package piscine

func LoafOfBread(str string) string {
	result := ""
	count := 0
	if len(str) < 5 && str != "" {
		result += "Invalid Output\n"
		return result
	}
	for i, ch := range str {
		if ch >= 'a' && ch <= 'z' && ch >= 'A' && ch <= 'Z' {
			if count == 5 {
				if i != (len(str) - 1) {
					result += string(' ')
					count = 0
				}
			} else {
				if !(ch == ' ') {
					count++
					result += string(ch)
				}
			}
		}
	}
	result += string('\n')
	return result
}
