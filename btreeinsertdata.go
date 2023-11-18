package piscine

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(root.Left.Data)
// 	fmt.Println(root.Data)
// 	fmt.Println(root.Right.Left.Data)
// 	fmt.Println(root.Right.Data)
// }

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		return newNode
	}

	if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	}

	return root
}
