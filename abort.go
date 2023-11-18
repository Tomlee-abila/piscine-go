package piscine

func Abort(a, b, c, d, e int) int {
	var array []int
	temp := 0
	array = append(array, a)
	array = append(array, b)
	array = append(array, c)
	array = append(array, d)
	array = append(array, e)
	for j := 0; j < 5; j++ {
		for i := 0; i < 4; i++ {
			if array[i] > array[i+1] {
				temp = array[i]
				array[i] = array[i+1]
				array[i+1] = temp
			}
		}
	}
	return array[2]
}
