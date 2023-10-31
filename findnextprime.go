package piscine

func FindNextPrime(nb int) int {
	for i := nb; i >= nb; i++ {
		if Prime(i) {
			return i
		}
	}
	return 1
}

func Prime(nb int) bool {
	if nb == 1 || nb < 0 || nb == 0 {
		return false
	}
	for i := 1; i < nb; i++ {
		if i != 1 && nb%i == 0 {
			return false
		}
	}
	return true
}
