package piscine

// package main

// import (
// 	"fmt"
// )

// type List = List2
// type Node = NodeL

// func PrintList(l *List) {
// 	link := l.Head
// 	for link != nil {
// 		fmt.Print(link.Data, " -> ")
// 		link = link.Next
// 	}
// 	fmt.Println(nil)
// }

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "I")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "something")
// 	ListPushBack(link, 2)

// 	fmt.Println("------list------")
// 	PrintList(link)
// 	ListClear(link)
// 	fmt.Println("------updated list------")
// 	PrintList(link)
// }

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List2 struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List2, data interface{}) {
// 	n := &NodeL{Data: data}

// 	if l.Head == nil {
// 		l.Head = n
// 	} else {
// 		l.Tail.Next = n
// 	}
// 	l.Tail = n
// }

func ListClear(l *List) {
	for l.Head != nil {
		temp := l.Head.Next
		l.Head = nil
		l.Head = temp
	}

	l.Tail = nil
}
