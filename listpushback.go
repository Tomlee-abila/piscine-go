package piscine

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	link := &List{}

// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, "man")
// 	ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}
