package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr[len(arr)/2]
}
