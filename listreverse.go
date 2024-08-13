package piscine

func ListReverse(l *List) {
	var next *NodeL
	var prev *NodeL
	current := l.Head
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
