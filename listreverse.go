package piscine

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushBack(link, 1)
// 	ListPushBack(link, 2)
// 	ListPushBack(link, 3)
// 	ListPushBack(link, 4)

// 	ListReverse(link)

// 	it := link.Head

// 	for it != nil {
// 		fmt.Println(it.Data)
// 		it = it.Next
// 	}

// 	fmt.Println("Tail", link.Tail)
// 	fmt.Println("Head", link.Head)
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}

// 	if Head == nil {
// 		Head = n
// 	} else {
// 		TaiNext = n
// 	}
// 	Tail = n
// }

func ListReverse(l *List) {
	current := l.Head
	var next *NodeL
	var prev *NodeL

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		next.Next = prev
		current = next
	}
	l.Head = prev
}
