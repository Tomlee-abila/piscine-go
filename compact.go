package piscine

func Compact(ptr *[]string) int {
	arr := *ptr
	count := 0
	for _, s := range arr {
		if s == "" {
			count++
		}
	}
	return count
}
