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
// 	fmt.Println(BTreeLevelCount(root))
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

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r1 := root
	r2 := root
	count1 := 0
	count2 := 0
	result := 0

	for r1 != nil {
		count1++
		r1 = r1.Right
	}

	for r2 != nil {
		count2++
		r2 = r2.Right
	}

	if count1 > count2 {
		result = count1
	} else {
		result = count2
	}

	return result
}
