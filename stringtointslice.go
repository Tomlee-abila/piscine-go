package piscine

func StringToIntSlice(str string) []int {
	var arr []int
	for _, ch := range str {
		arr = append(arr, int(ch))
	}
	return arr
}
