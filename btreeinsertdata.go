package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node with the given data into the binary search tree
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		// Create a new node if the tree is empty
		return &TreeNode{Data: data}
	}

	// Recursive function to find the correct insertion point
	var insert func(node *TreeNode, data string) *TreeNode
	insert = func(node *TreeNode, data string) *TreeNode {
		if data < node.Data {
			// Insert into the left subtree
			if node.Left == nil {
				node.Left = &TreeNode{Data: data, Parent: node}
				return node.Left
			}
			return insert(node.Left, data)
		} else if data > node.Data {
			// Insert into the right subtree
			if node.Right == nil {
				node.Right = &TreeNode{Data: data, Parent: node}
				return node.Right
			}
			return insert(node.Right, data)
		}
		// Data is equal to the current node's data; do nothing to avoid duplicates
		return nil
	}

	return insert(root, data)
}
