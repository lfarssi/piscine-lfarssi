package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node with the given data into the binary search tree
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		// If the tree is empty, create a new node and return it as the root
		return &TreeNode{Data: data}
	}

	// Helper function to insert data recursively
	var insert func(node *TreeNode, data string) *TreeNode
	insert = func(node *TreeNode, data string) *TreeNode {
		if data < node.Data {
			if node.Left == nil {
				// Insert the new node as the left child
				node.Left = &TreeNode{Data: data, Parent: node}
				return node.Left
			}
			return insert(node.Left, data)
		} else if data > node.Data {
			if node.Right == nil {
				// Insert the new node as the right child
				node.Right = &TreeNode{Data: data, Parent: node}
				return node.Right
			}
			return insert(node.Right, data)
		}
		// If data == node.Data, we do nothing as it's a duplicate
		return nil
	}

	return insert(root, data)
}
