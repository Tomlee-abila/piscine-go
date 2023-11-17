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

// 	if l.Head == nil {
// 		l.Head = n
// 	} else {
// 		l.Tail.Next = n
// 	}
// 	l.Tail = n
// }

// func ListPushFront(l *List, data interface{}) {
// 	newerNode := &NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = newerNode
// 	} else {
// 		newerNode.Next = l.Head
// 		l.Head = newerNode
// 	}
// }

func ListReverse(l *List) {
	link := &List{}

	for l.Head != nil {
		ListPushFront(link, l.Head.Data)
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
	}

	for link.Head != nil {
		ListPushBack(l, link.Head.Data)
		link.Head = link.Head.Next
	}
	// current := l.Head
	// var next *NodeL
	// var prev *NodeL = &NodeL{}

	// for current != nil {
	// 	next = current.Next
	// 	current.Next = prev
	// 	prev = current
	// 	current = next
	// 	// next.Next = prev
	// }
	// l.Head = prev

}
