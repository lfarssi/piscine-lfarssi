package piscine

func DescendAppendRange(max, min int) []int {
	var arr []int
	if max <= min {
		return nil
	}
	for i := max; i >= min; i-- {
		arr = append(arr, i)
	}
	return arr
}
