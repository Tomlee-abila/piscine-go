package piscine

func ReverseMenuIndex(menu []string) []string {
	arrayMenu := make([]string, len(menu))
	for i := len(menu) - 1; i >= 0; i-- {
		arrayMenu[(len(menu)-1)-i] = menu[i]
	}
	return arrayMenu
}
