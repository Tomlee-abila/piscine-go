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
// 	fmt.Println(BTreeIsBinary(root))
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
// 		root.Right.Parent = root
// 	}

// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 		root.Left.Parent = root
// 	}

// 	return root
// }

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// If both left and right subtrees are non-nil
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}

	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}

	// Recursively check both left and right subtrees
	if !(BTreeIsBinary(root.Left)) || !(BTreeIsBinary(root.Right)) {
		return false
	}

	return true
}
