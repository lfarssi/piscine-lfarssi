package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	// Remove nodes from the head
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Remove nodes in the middle
	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
			if current.Next == nil {
				l.Tail = current
			}
		} else {
			current = current.Next
		}
	}
}
