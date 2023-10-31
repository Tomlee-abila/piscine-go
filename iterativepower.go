package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 || power < 0 {
		return 0
	} else {
		result := nb
		for i := 1; i < power; i++ {
			result = result * nb
			if result < 0 {
				return 0
			}
		}
		return result
	}
}
