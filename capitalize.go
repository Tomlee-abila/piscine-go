package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	wordStart := true

	for i, r := range runes {
		if isAlphaNumeric(r) {
			if wordStart {
				runes[i] = toUpper(r)
				wordStart = false
			} else {
				runes[i] = toLower(r)
			}
		} else {
			wordStart = true
		}
	}

	return string(runes)
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}
