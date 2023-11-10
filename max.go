package piscine

func Max(a []int) int {
	result := 0

	for _, num := range a {
		if num > result {
			result = num
		}
	}
	return result
}
