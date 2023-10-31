package piscine

func FindNextPrime(nb int) int {
	result := 0
	for i := nb; i >= nb; i++ {
		if Prime(i) {
			result = i
			return i
		}
	}
	return result
}

func Prime(nb int) bool {
	if nb < 2 {
		return true
	}
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 11 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 || nb%5 == 0 || nb%7 == 0 || nb%11 == 0 {
		return false
	}
	for i := 13; i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
