package piscine

func Unmatch(a []int) int {
	Double := false
	count := 0
	var result int
	for i, n := range a {
		Double = false
		count = 0
		for j, m := range a {
			count++
			if i != j {
				if n == m {
					Double = true
				}
			}
		}
		if count == len(a) && Double == false {
			result = n
		}
	}
	return result
}
