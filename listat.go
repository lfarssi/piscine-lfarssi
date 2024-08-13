package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	i := 0
	for l != nil && i < pos {
		l = l.Next
		i++
	}
	return l
}
