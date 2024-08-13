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
  current := l.Head
  for current.Next != nil {
    if current.Data == data_ref {
      current.Data = nil
    }
    current = current.Next
  }
  return current.Data
}