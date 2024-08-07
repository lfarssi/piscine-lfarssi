package piscine

var down int

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start /= 2
		} else if start%2 == 1 {
			start = (start * 3) + 1
		}
		down++
	}
	return down
}
