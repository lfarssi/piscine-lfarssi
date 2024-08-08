package piscine

func PodiumPosition(podium [][]string) [][]string {
	var arr [][]string
	for i := len(podium) - 1; i >= 0; i-- {
		arr = append(arr, podium[i])
	}
	return arr
}
