package piscine

// package main

// import (
// 	"fmt"
// )

// func PrintList(l *List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}

// 	fmt.Print(nil, "\n")
// }

// func main() {
// 	link := &List{}
// 	link2 := &List{}

// 	fmt.Println("----normal state----")
// 	ListPushBack(link2, 1)
// 	PrintList(link2)
// 	ListRemoveIf(link2, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link2)
// 	fmt.Println()

// 	fmt.Println("----normal state----")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "There")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "How")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)
// 	PrintList(link)

// 	ListRemoveIf(link, 1)
// 	fmt.Println("------answer-----")
// 	PrintList(link)
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

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	// if l.Head == nil {
	// 	return
	// }
	// if l.Head.Data == data_ref {
	// 	l.Head = l.Head.Next
	// }
	// var prev *NodeL
	// curr := l.Head

	// for curr != nil {
	// 	if curr.Data == data_ref {
	// 		prev.Next = curr.Next
	// 	} else {
	// 		prev = curr
	// 	}
	// 	curr = curr.Next
	// }
	// if prev != nil {
	// 	l.Tail = prev
	// }

	if l.Head == nil {
		return
	}
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	current := l.Head
	var pre *NodeL
	for current != nil {
		if data_ref == current.Data {
			pre.Next = current.Next
		} else {
			pre = current
		}
		current = current.Next
	}

	if pre != nil {
		l.Tail = pre
	}
}
