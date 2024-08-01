package piscine

func ConcatParams(args []string) string {
	var r string
	for _, j := range args {
		r = r + "\n" + j
	}
	r = r[1:]
	return r
}
