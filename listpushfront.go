package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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
