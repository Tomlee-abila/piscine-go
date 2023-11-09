package piscine

func Unmatch(a []int) int {
	Double := 0
	count := 0
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
			return n
		}
	}
	return -1
}
