package piscine

func Compact(ptr *[]string) int {
	arr := *ptr
	c := 0
	for _, i := range arr {
		if i != "" {
			arr[c] = i
			c++
		}
	}
	*ptr = arr[:c]
	return c
}
