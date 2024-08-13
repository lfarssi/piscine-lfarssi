package piscine

func ListLast(l *List) interface{} {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
	return l.Head
}
