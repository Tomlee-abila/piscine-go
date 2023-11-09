package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	arrDeck := deck
	player := 1

	for i := 0; i < 12; i += 3 {
		fmt.Println("Player", player, ":", arrDeck[i], ",", arrDeck[i+1], ",", arrDeck[i+2])
		player++
	}
}
