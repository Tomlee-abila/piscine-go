package piscine

// package main

// import (
// 	"fmt"
// )

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
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

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	lNode := l
	if pos == 0 && lNode != nil {
		return l
	}
	if l == nil {
		return nil
	}

	for lNode != nil {
		if count == pos {
			return lNode
		}
		lNode = lNode.Next
		count++
	}
	return lNode
}
