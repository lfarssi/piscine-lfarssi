package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeDeleteNode deletes a node from the tree
func BTreeDeleteNode(root, nodeToRemove *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Traverse the tree to find the node to remove
	if nodeToRemove.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, nodeToRemove)
	} else if nodeToRemove.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, nodeToRemove)
	} else {
		// Node with only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Node with two children
		successor := BTreeMin(root.Right)
		root.Data = successor.Data
		root.Right = BTreeDeleteNode(root.Right, successor)
	}

	return root
}

// BTreeMin finds the node with the minimum value in the tree
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
