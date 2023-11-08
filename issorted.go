package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	x := 0
	y := 0
	for i := 0; i < (len(a) - 1); i++ {
		if !((f(a[i], a[i+1])) >= 0) {
			x = 1
		}
	}

	for i := 0; i < (len(a) - 1); i++ {
		if !((f(a[i], a[i+1])) <= 0) {
			y = 1
		}
	}

	if x == 0 || y == 0 {
		return true
	}
	return false
}
