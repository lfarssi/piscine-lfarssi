package piscine

func ListMerge(l1 *List, l2 *List) {
	// If l1 is empty
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// If l2 is empty, no need to merge anything
	if l2.Head == nil {
		return
	}

	// Find the end of l1
	current := l1.Head
	for current.Next != nil {
		current = current.Next
	}

	// Link the end of l1 to the start of l2
	current.Next = l2.Head

	// Update the tail of l1 to be the tail of l2
	l1.Tail = l2.Tail
}
