package piscine

func Sqrt(nb int) int {
	sqr := 0
	for i := 1; i <= nb; i++ {
		sqr = i * i
		if sqr == nb {
			return i
		} else if i == nb {
			return 0
		}
	}
	return 0
}
