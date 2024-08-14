package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

// ListSort sorts the linked list using merge sort algorithm
func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	// Step 1: Split the list into two halves
	middle := getMiddle(l)
	nextToMiddle := middle.Next
	middle.Next = nil

	// Step 2: Recursively sort the two halves
	left := ListSort(l)
	right := ListSort(nextToMiddle)

	// Step 3: Merge the sorted halves
	return merge(left, right)
}

// getMiddle finds the middle of the linked list
func getMiddle(head *NodeI) *NodeI {
	if head == nil {
		return head
	}

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// merge merges two sorted linked lists
func merge(left *NodeI, right *NodeI) *NodeI {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	var result *NodeI
	if left.Data <= right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}
