package piscine

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(ListSize(link))
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
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

func ListSize(l *List) int {
	length := 0

	for l.Head != nil {
		l.Head = l.Head.Next
		length++
	}
	return length
}
