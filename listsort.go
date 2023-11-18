package piscine

// package main

import (
	"fmt"
)

// func PrintList(l *NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *NodeI, data int) *NodeI {
// 	n := &NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

// func main() {
// 	var link *NodeI

// 	link = listPushBack(link, 5)
// 	link = listPushBack(link, 4)
// 	link = listPushBack(link, 3)
// 	link = listPushBack(link, 2)
// 	link = listPushBack(link, 1)

// 	PrintList(ListSort(link))
// }

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	head := l
	liNode := l

	if head.Data > head.Next.Data {
		temp := head
		fmt.Println("head.Data =", head.Data, "head.Next.Data =", head.Next.Data)
		head = head.Next
		head.Next = temp
		// head.Next = head.Next.Next
		// temp2 := head.Next
		fmt.Println("head.Data =", head.Data, "head.Next.Data =", head.Next.Data)
		// head.Next.Next = head.Next
		// head = temp
		// head.Next = temp2
		fmt.Println("head.Data =", head.Data, "head.Next.Data =", head.Next.Data)
	}
	return liNode
}
