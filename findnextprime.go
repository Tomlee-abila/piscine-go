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
		return false
	}
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 11 || nb == 13 || nb == 17 || nb == 19 || nb == 23 || nb == 29 || nb == 31 || nb == 37 || nb == 41 || nb == 43 || nb == 47 || nb == 53 || nb == 59 || nb == 61 || nb == 67 || nb == 71 || nb == 73 || nb == 79 || nb == 83 || nb == 89 || nb == 97 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 || nb%5 == 0 || nb%7 == 0 || nb%11 == 0 || nb%13 || nb%17 || nb%19 || nb%23 || nb%29 || nb%31 || nb%37 || nb%41 || nb%43 || nb%47 || nb%53 || nb%59 || nb%61 || nb%67 || nb%71 || nb%73 || nb%79 || nb%83 || nb%89 || nb%97 {
		return false
	}
	for i := 97; i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
