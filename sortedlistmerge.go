package piscine

// SortedListMerge merges two sorted linked lists into one sorted linked list
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Create a dummy node to act as the start of the merged list
	dummy := &NodeI{}
	current := dummy

	// Traverse both lists and append the smaller node to the merged list
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			current.Next = n1
			n1 = n1.Next
		} else {
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Append the remaining nodes from either list
	if n1 != nil {
		current.Next = n1
	}
	if n2 != nil {
		current.Next = n2
	}

	// Return the next node of the dummy node, which is the head of the merged list
	return dummy.Next
}
