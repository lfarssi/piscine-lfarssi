package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	if current.Next != nil {
		current = current.Next
	}
	return current.Data
}
