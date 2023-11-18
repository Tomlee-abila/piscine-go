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
// 	BTreeApplyPreorder(root, fmt.Println)
// }

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	newNode := &TreeNode{Data: data}

// 	if root == nil {
// 		return newNode
// 	}

// 	if data > root.Data {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	}

// 	return root
// }

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	f(root.Data)
	BTreeApplyPreorder(root.Left, f)
	BTreeApplyPreorder(root.Right, f)
}
