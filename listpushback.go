package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	newN := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newN
		l.Tail = newN
	} else {
		l.Tail.Next = newN
		l.Tail = newN
	}
}
