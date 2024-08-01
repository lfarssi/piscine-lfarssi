package piscine

func ConcatParams(args []string) string {
	var r string
	for _, j := range args {
		r = r + j + "\n"
	}
	r = r[1:]
	return r
}
