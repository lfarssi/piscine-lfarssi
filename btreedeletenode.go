package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

// BTreeDeleteNode deletes the node from the tree rooted at `root` and returns the new root of the tree
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil && node.Right == nil {
		// Node is a leaf
		if node.Parent == nil {
			return nil // Tree was a single node
		}
		if node == node.Parent.Left {
			node.Parent.Left = nil
		} else {
			node.Parent.Right = nil
		}
	} else if node.Left == nil {
		// Node has only right child
		if node.Parent == nil {
			return node.Right // New root of the tree
		}
		if node == node.Parent.Left {
			node.Parent.Left = node.Right
		} else {
			node.Parent.Right = node.Right
		}
		node.Right.Parent = node.Parent
	} else if node.Right == nil {
		// Node has only left child
		if node.Parent == nil {
			return node.Left // New root of the tree
		}
		if node == node.Parent.Left {
			node.Parent.Left = node.Left
		} else {
			node.Parent.Right = node.Left
		}
		node.Left.Parent = node.Parent
	} else {
		// Node has two children
		successor := minimum(node.Right)
		node.Data = successor.Data
		root = BTreeDeleteNode(root, successor)
	}

	return root
}

// minimum returns the node with the minimum value in the subtree rooted at `node`
func minimum(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
