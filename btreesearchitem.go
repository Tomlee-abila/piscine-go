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
// 	selected := BTreeSearchItem(root, "7")
// 	fmt.Print("Item selected -> ")
// 	if selected != nil {
// 		fmt.Println(selected.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Parent of selected item -> ")
// 	if selected.Parent != nil {
// 		fmt.Println(selected.Parent.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Left child of selected item -> ")
// 	if selected.Left != nil {
// 		fmt.Println(selected.Left.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}

// 	fmt.Print("Right child of selected item -> ")
// 	if selected.Right != nil {
// 		fmt.Println(selected.Right.Data)
// 	} else {
// 		fmt.Println("nil")
// 	}
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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	var result *TreeNode

	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}
	if elem > root.Data {
		result = BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		result = BTreeSearchItem(root.Left, elem)
	}
	return result
}
