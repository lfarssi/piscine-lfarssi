package piscine

func ListPushFront(l *List, data interface{}) {
	newN := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newN
		l.Tail = newN
	} else {
		newN.Next = l.Head
		l.Head = newN
	}
}
