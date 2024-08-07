package piscine

func StringToIntSlice(str string) []int {
	var arr []int
	for _, i := range str {
		arr = append(arr, int(i))
	}
	return arr
}
