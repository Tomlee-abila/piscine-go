package piscine

func ShoppingListSort(slice []string) []string {
	array := slice
	for j := 0; j < len(array); j++ {
		for i := 0; i < len(array)-1; i++ {
			if len(array[i]) > len(array[i+1]) {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp

			}
		}
	}
	return array
}
