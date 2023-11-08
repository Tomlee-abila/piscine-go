package piscine

// func PrintNbr(n int) {
// 	fmt.Print(n)

// }
func ForEach(f func(int), a []int) {
	for _, num := range a {
		f(num)
	}
	fmt.Println()
}
