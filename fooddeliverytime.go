package piscine

type food struct {
	burger  int
	chips   int
	nuggets int
}

func setPoint(ptr *food) {
	ptr.burger = 15
	ptr.chips = 10
	ptr.nuggets = 12
}

func FoodDeliveryTime(order string) int {
	preptime := &food{}
	setPoint(preptime)
	runeOrder := []rune(order)
	totalTime := 0

	for i, ch := range runeOrder {
		if ch == 'b' {
			for j, rn := range "burger" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			totalTime += preptime.burger
		} else if ch == 'c' {
			for j, rn := range "chips" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			totalTime += preptime.chips
		} else if ch == 'n' {
			for j, rn := range "nuggets" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			totalTime += preptime.nuggets
		}
	}
	return totalTime
}
