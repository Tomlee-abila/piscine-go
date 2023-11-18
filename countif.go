package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, num := range tab {
		if f(num) {
			count++
		}
	}
	return count
}
