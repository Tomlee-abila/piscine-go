package piscine

func Unmatch(a []int) int {
	Double := 0
	count := 0
	var result int
	for i, n := range a {
		Double = 0
		count = 0
		for j, m := range a {
			count++
			if i != j {
				if n == m {
					Double++
				}
			}
		}
		if count == len(a) && Double%2 == 0 {
			result = n
		}
	}
	return result
}
