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
	ch := 0
	br := 0
	nu := 0

	for i, ch := range runeOrder {
		if ch == 'b' {
			for j, rn := range "burger" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			br++
			totalTime += preptime.burger
		} else if ch == 'c' {
			for j, rn := range "chips" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			ch++
			totalTime += preptime.chips
		} else if ch == 'n' {
			for j, rn := range "nuggets" {
				if rn != runeOrder[i+j] {
					return 404
				}
			}
			nu++
			totalTime += preptime.nuggets
		}
	}
	if (ch*len("chips"))+(br*len("burger"))+(nu*len("nuggets")) != len(runeOrder) {
		return 404
	}
	return totalTime
}
