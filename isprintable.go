package piscine

func IsPrintable(s string) bool {
	for _, str := range s {
		if !(alphaNum(str)) {
			return false
		}
	}
	return true
}
