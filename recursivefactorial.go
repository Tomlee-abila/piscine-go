package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 39 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := nb * RecursiveFactorial(nb-1)
		if result < 0 {
			return 0
		} else {
			return result
		}
	}
}
