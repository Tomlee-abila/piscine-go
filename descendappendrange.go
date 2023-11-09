package piscine

func DescendAppendRange(max, min int) []int {
	var result []int
	if max > min {
		for i := min + 1; i <= max; i++ {
			result = append(result, i)
		}
	} else {
		result = []int{}
	}
	return result
}
