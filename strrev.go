package piscine

func StrRev(s string) string {
	aString := []byte(s)
	bString := []byte(s)

	for i := len(s) - 1; i >= 0; i-- {
		aString[len(s)-i-1] = bString[i]
	}

	s = string(aString)
	return s
}
