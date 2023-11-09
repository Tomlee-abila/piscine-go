package piscine

func splitWords(s string) []string {
	word := ""
	var result []string
	firstWord := true
	for i, chr := range s {
		if chr >= 'a' && chr <= 'z' || chr >= 'A' && chr <= 'Z' {
			word += string(chr)
		}
		if chr == ' ' || i == len(s)-1 {
			firstWord = false
		}
		if firstWord == false && word != "" {
			firstWord = true
			result = append(result, word)
			word = ""
		}
	}
	return result
}
