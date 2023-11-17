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

type List2 struct {
	Head *NodeL
	Tail *NodeL
	prev *NodeL
	next *NodeL
}

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = n
// 	} else {
// 		l.Tail.Next = n
// 	}
// 	l.Tail = n
// }

func ListReverse(l *List2) {
	current := l.Head

	for current != nil {
		l.next = current.Next
		current.Next = l.prev
		l.prev = current
		l.next.Next = l.prev
		current = l.next
	}
	l.Head = l.prev
}
