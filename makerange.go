package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	Range := max - min
	result := make([]int, Range)

	for i := 0; i < Range; i++ {
		result[i] = min + i
	}
	return result
}
