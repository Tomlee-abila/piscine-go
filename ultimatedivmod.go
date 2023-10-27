package piscine

func UltimateDivMod(a *int, b *int) {
	*b = *a % *b
	*a = *a / (*b + *b)
}
