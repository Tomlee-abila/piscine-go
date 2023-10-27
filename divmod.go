package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*mod = b - 1
	*div = (a - 1) / 2
}
