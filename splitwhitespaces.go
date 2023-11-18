package piscine

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var words []string
	wordStart := -1

	for i, r := range s {
		if isWhitespace(r) {
			if wordStart != -1 {
				words = append(words, s[wordStart:i])
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
	}

	if wordStart != -1 {
		words = append(words, s[wordStart:])
	}

	return words
}
