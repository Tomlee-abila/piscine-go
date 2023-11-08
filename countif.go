package piscine

func CountIf(f func(string) bool, tab []string) int {
	for i, num := range tab {
		if f(num) {
			return i + 1
		}
	}
	return 0
}
