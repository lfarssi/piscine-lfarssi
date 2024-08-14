package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts a new node with the given data into the binary search tree
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Helper function to insert data recursively
	var insert func(node *TreeNode, data string) *TreeNode
	insert = func(node *TreeNode, data string) *TreeNode {
		if data < node.Data {
			if node.Left == nil {
				node.Left = &TreeNode{Data: data, Parent: node}
				return node.Left
			}
			return insert(node.Left, data)
		} else if data > node.Data {
			if node.Right == nil {
				node.Right = &TreeNode{Data: data, Parent: node}
				return node.Right
			}
			return insert(node.Right, data)
		}
		return nil
	}

	return insert(root, data)
}
