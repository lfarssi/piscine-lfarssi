package main

import (
	"fmt"
	"piscine"
)

func main() {
	// Create the root node
	root := &piscine.TreeNode{Data: "08"}

	// Insert nodes
	piscine.BTreeInsertData(root, "x")
	piscine.BTreeInsertData(root, "y")

	// Print the structure of the tree
	printTree(root)
}

// Helper function to print the tree structure
func printTree(node *piscine.TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		fmt.Println(node.Left.Data)
	}
	fmt.Println(node.Data)
	if node.Right != nil {
		fmt.Println(node.Right.Data)
	}
	printTree(node.Left)
	printTree(node.Right)
}
