package piscine

func ListPushFront(l *List, data interface{}) {
	newN := &NodeL{Data: data}
	if l.Head != nil {
		l.Head = newN
		l.Tail = newN
	} else {
		l.Tail.Next = newN
		l.Tail = newN
	}
}
