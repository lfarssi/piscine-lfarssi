package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return nil
	}
	var arr []int
	for i := max; i >= min; i-- {
		arr = append(arr, i)
	}
	return arr
}
