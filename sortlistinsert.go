package piscine

// SortListInsert inserts a new node with data_ref into the sorted linked list
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	// Create the new node
	newNode := &NodeI{Data: data_ref}

	// Case 1: Empty list or new node should be the new head
	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	// Case 2: Insert in the middle or end
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Insert the new node
	newNode.Next = current.Next
	current.Next = newNode

	return l
}
