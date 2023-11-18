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
// 	max := BTreeMax(root)
// 	fmt.Println(max.Data)
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

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	result := root

	// node = BTreeMax(root.Right)
	for node != nil {
		result = node
		node = node.Right
	}
	return result
}
