package piscine

func ReverseMenuIndex(menu []string) []string {
	var arr []string
	for i := len(menu) - 1; i >= 0; i-- {
		arr = append(arr, menu[i])
	}
	return arr
}
