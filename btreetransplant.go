package piscine

// BTreeTransplant replaces the subtree rooted at `node` with the subtree rooted at `rplc` in the tree rooted at `root`
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Parent == nil {
		// Replacing the root
		if rplc != nil {
			rplc.Parent = nil
		}
		return rplc
	}

	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}

	return root
}
