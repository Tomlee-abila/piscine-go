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
// 	node := BTreeSearchItem(root, "1")
// 	rplc := &TreeNode{Data: "3"}
// 	root = BTreeTransplant(root, node, rplc)
// 	BTreeApplyInorder(root, fmt.Println)
// }

// type TreeNode struct {
// 	Left, Right, Parent *TreeNode
// 	Data                string
// }

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	var result *TreeNode

// 	if root == nil {
// 		return nil
// 	}

// 	if root.Data == elem {
// 		return root
// 	}
// 	if elem > root.Data {
// 		result = BTreeSearchItem(root.Right, elem)
// 	} else if elem < root.Data {
// 		result = BTreeSearchItem(root.Left, elem)
// 	}
// 	return result
// }

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}
// 	BTreeApplyInorder(root.Left, f)
// 	f(root.Data)
// 	BTreeApplyInorder(root.Right, f)
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

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// If the node to be replaced is the root of the tree
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// If the node is the left child of its parent
		node.Parent.Left = rplc
	} else {
		// If the node is the right child of its parent
		node.Parent.Right = rplc
	}

	// If the replacement node is not nil, update its parent
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	// Disconnect the replaced node from the tree
	node.Parent = nil
	node.Left = nil
	node.Right = nil

	return root
}
