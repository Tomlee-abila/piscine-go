package piscine

func Compact(ptr *[]string) int {
	arr := *ptr
	var arr1 []string
	count := 0
	for _, s := range arr {
		if !(s == "") {
			count++
			arr1 = append(arr1, s)
		}
	}
	*ptr = arr1
	return count
}
