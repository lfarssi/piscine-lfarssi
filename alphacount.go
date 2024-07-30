package piscine

func AlphaCount(s string) int {
	count := 0
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			count++
		}	
	}
	return count
}
