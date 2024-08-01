package piscine

func ConcatParams(args []string) string {
	arr := make([]string, len(args))
	for _, i := range args {
		arr = append(arr, i)
	}
	r := ""
	for _, j := range arr {
		r = j
	}
	return r
}
