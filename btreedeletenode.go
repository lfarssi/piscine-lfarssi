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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}

func BTreeIsBinary(root *TreeNode) bool {
	return isBinary(root, nil, nil)
}

func isBinary(node *TreeNode, min *TreeNode, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && node.Data <= min.Data {
		return false
	}
	if max != nil && node.Data >= max.Data {
		return false
	}
	return isBinary(node.Left, min, node) && isBinary(node.Right, node, max)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}

