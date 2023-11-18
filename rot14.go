package piscine

func Rot14(t string) string {
	s := []rune(t)
	for i, r := range s {
		if r >= 'a' && r <= 'z' {
			s[i] += 14
			if s[i] > 'z' {
				s[i] = s[i] - 'z' + 'a' - 1
			}
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] += 14
			if s[i] > 'Z' {
				s[i] = s[i] - 'Z' + 'A' - 1
			}
		}
	}
	return string(s)
}
