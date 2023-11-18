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
// 	BTreeApplyByLevel(root, fmt.Println)
// }

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	f(root.Data)
	BTreeApplyByLevel(root.Left, f)
	BTreeApplyByLevel(root.Right, f)
}
